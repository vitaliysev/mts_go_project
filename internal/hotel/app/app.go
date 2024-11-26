package app

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vitaliysev/mts_go_project/internal/hotel/api/hotel"
	"github.com/vitaliysev/mts_go_project/internal/hotel/closer"
	"github.com/vitaliysev/mts_go_project/internal/hotel/config"
	"github.com/vitaliysev/mts_go_project/internal/hotel/interceptor"
	desc "github.com/vitaliysev/mts_go_project/pkg/hotel_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"sync"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	restServer      *http.Server
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
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := a.runRESTServer()
		if err != nil {
			log.Fatalf("failed to run REST server: %v", err)
		}
	}()
	go func() {
		defer wg.Done()
		err := a.runGRPCServer()
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
		a.initGRPCServer,
		a.initRESTServer,
	}
	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
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

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
	)

	reflection.Register(a.grpcServer)

	desc.RegisterHotelV1Server(a.grpcServer, a.serviceProvider.HotelImpl(ctx))

	return nil
}
func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initRESTServer(ctx context.Context) error {
	log1 := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Post("/saveHotel", hotel.NewSave(ctx, log1, a.serviceProvider.hotelImpl))
	router.Put("/updateHotel", hotel.NewUpdate(ctx, log1, a.serviceProvider.hotelImpl))
	router.Get("/getHotel", hotel.NewGetHotel(ctx, log1, a.serviceProvider.hotelImpl))
	router.Get("/getHotels", hotel.NewGetHotels(ctx, log1, a.serviceProvider.hotelImpl))
	server := &http.Server{
		Addr:    a.serviceProvider.RESTConfig().Address(),
		Handler: router,
	}
	a.restServer = server
	return nil
}

func (a *App) runRESTServer() error {
	log.Printf("REST server is running on %s", a.serviceProvider.RESTConfig().Address())
	if err := a.restServer.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
