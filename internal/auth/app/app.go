package app

import (
	"context"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/vitaliysev/mts_go_project/internal/auth/closer"
	"github.com/vitaliysev/mts_go_project/internal/auth/config"
	interceptor2 "github.com/vitaliysev/mts_go_project/internal/auth/interceptor"
	"github.com/vitaliysev/mts_go_project/internal/auth/logger/interceptor"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
	"google.golang.org/grpc/credentials/insecure"

	//	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
	"net"
	"net/http"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/vitaliysev/mts_go_project/internal/auth/logger"
	descAuth "github.com/vitaliysev/mts_go_project/pkg/auth_v1"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()

		err := a.runGRPCAuthServer()
		if err != nil {
			log.Fatalf("failed to run GRPC server: %v", err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCAuthServer,
		a.initTracing,
	}
	logger.Init(getCore(getAtomicLevel()))
	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
func (a *App) initTracing(ctx context.Context) error {
	err := tracing.NewTracer("http://localhost:14268/api/traces", "Auth-service")
	return err
}
func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCAuthServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(interceptor.LogInterceptor,
			interceptor2.ServerTracingInterceptor)),
	)

	reflection.Register(a.grpcServer)

	descAuth.RegisterAuthV1Server(a.grpcServer, a.serviceProvider.GRPCAuthImpl(ctx))

	return nil
}

func (a *App) runGRPCAuthServer() error {
	log.Printf("GRPC auth server is running on %s", a.serviceProvider.GRPCAuthConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCAuthConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
