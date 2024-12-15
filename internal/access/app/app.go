package app

import (
	"context"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/vitaliysev/mts_go_project/internal/access/closer"
	"github.com/vitaliysev/mts_go_project/internal/access/config"
	"github.com/vitaliysev/mts_go_project/internal/access/interceptor"
	"github.com/vitaliysev/mts_go_project/internal/tracing"

	//	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
	"net"
	"net/http"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	descAccess "github.com/vitaliysev/mts_go_project/pkg/access_v1"
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

		err := a.runGRPCAccessServer()
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
		a.initGRPCAccessServer,
		a.initTracing,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
func (a *App) initTracing(ctx context.Context) error {
	err := tracing.NewTracer("http://localhost:14268/api/traces", "Access-service")
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

func (a *App) initGRPCAccessServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()), grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
		interceptor.ServerTracingInterceptor)))

	reflection.Register(a.grpcServer)

	descAccess.RegisterAccessV1Server(a.grpcServer, a.serviceProvider.GRPCAccessImpl(ctx))

	return nil
}

func (a *App) runGRPCAccessServer() error {
	log.Printf("GRPC access access server is running on %s", a.serviceProvider.GRPCAccessConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCAccessConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
