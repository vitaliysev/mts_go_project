package app

import (
	"context"
	"github.com/go-chi/chi/v5"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/natefinch/lumberjack"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vitaliysev/mts_go_project/internal/hotel/api/hotel"
	"github.com/vitaliysev/mts_go_project/internal/hotel/closer"
	"github.com/vitaliysev/mts_go_project/internal/hotel/config"
	"github.com/vitaliysev/mts_go_project/internal/hotel/interceptor"
	"github.com/vitaliysev/mts_go_project/internal/hotel/metrics"
	"github.com/vitaliysev/mts_go_project/internal/lib/logger"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
	desc "github.com/vitaliysev/mts_go_project/pkg/hotel_v1"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
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
	wg.Add(3)
	go func() {
		err := runPrometheus()
		if err != nil {
			log.Fatal(err)
		}
	}()
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
func runPrometheus() error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	prometheusServer := &http.Server{
		Addr:    "localhost:2112",
		Handler: mux,
	}

	log.Printf("Prometheus server is running on %s", "localhost:2112")

	err := prometheusServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
		a.initRESTServer,
		a.initLogger,
		a.initTracing,
		a.initMetrics,
	}
	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
func (a *App) initMetrics(ctx context.Context) error {
	err := metric.Init(ctx)
	return err
}
func (a *App) initLogger(ctx context.Context) error {
	logger.Init(getCore(getAtomicLevel()))
	return nil
}
func (a *App) initTracing(ctx context.Context) error {
	err := tracing.NewTracer("http://localhost:14268/api/traces", "Hotel-service")
	return err
}
func getCore(level zap.AtomicLevel) zapcore.Core {
	stdout := zapcore.AddSync(os.Stdout)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     7, // days
	})

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	return zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, file, level),
	)
}
func getAtomicLevel() zap.AtomicLevel {
	logLevel := os.Getenv("LOG_LEVEL")
	var level zapcore.Level
	if err := level.Set(logLevel); err != nil {
		log.Fatalf("failed to set log level: %v", err)
	}

	return zap.NewAtomicLevelAt(level)
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
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(interceptor.MetricsInterceptor, interceptor.LogInterceptor,
			interceptor.ServerTracingInterceptor, interceptor.ValidateInterceptor)),
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
	router := chi.NewRouter()
	router.Post("/saveHotel", hotel.NewSave(ctx, a.serviceProvider.hotelImpl))
	router.Put("/updateHotel", hotel.NewUpdate(ctx))
	router.Get("/getHotel", hotel.NewGetHotel(ctx, a.serviceProvider.hotelImpl))
	router.Get("/getHotels", hotel.NewGetHotels(ctx, a.serviceProvider.hotelImpl))
	a.restServer = &http.Server{
		Addr:    a.serviceProvider.RESTConfig().Address(),
		Handler: router,
	}
	return nil
}

func (a *App) runRESTServer() error {
	log.Printf("REST server is running on %s", a.serviceProvider.restConfig.Address())
	if err := a.restServer.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
