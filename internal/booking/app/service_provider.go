package app

import (
	"context"
	"flag"
	"github.com/natefinch/lumberjack"
	"github.com/vitaliysev/mts_go_project/internal/booking/api/booking_grpc"
	"github.com/vitaliysev/mts_go_project/internal/booking/api/booking_http"
	"github.com/vitaliysev/mts_go_project/internal/booking/client/db"
	"github.com/vitaliysev/mts_go_project/internal/booking/client/db/pg"
	"github.com/vitaliysev/mts_go_project/internal/booking/client/db/transaction"
	"github.com/vitaliysev/mts_go_project/internal/booking/closer"
	config2 "github.com/vitaliysev/mts_go_project/internal/booking/config"
	"github.com/vitaliysev/mts_go_project/internal/booking/repository"
	bookRepository "github.com/vitaliysev/mts_go_project/internal/booking/repository/booking"
	"github.com/vitaliysev/mts_go_project/internal/booking/service"
	bookService "github.com/vitaliysev/mts_go_project/internal/booking/service/booking"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var logLevel = flag.String("l", "info", "log level")

type serviceProvider struct {
	pgConfig   config2.PGConfig
	grpcConfig config2.GRPCConfig
	httpConfig config2.HTTPConfig

	dbClient       db.Client
	txManager      db.TxManager
	bookRepository repository.BookRepository

	bookService service.BookService

	bookgrpcImpl *booking_grpc.Implementation
	bookhttpImpl *booking_http.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config2.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config2.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config2.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config2.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config2.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config2.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) BookingRepository(ctx context.Context) repository.BookRepository {
	if s.bookRepository == nil {
		s.bookRepository = bookRepository.NewRepository(s.DBClient(ctx))
	}

	return s.bookRepository
}

func (s *serviceProvider) BookingService(ctx context.Context) service.BookService {
	if s.bookService == nil {
		s.bookService = bookService.NewService(
			s.BookingRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.bookService
}

func (s *serviceProvider) GRPCBookingImpl(ctx context.Context) *booking_grpc.Implementation {
	if s.bookgrpcImpl == nil {
		s.bookgrpcImpl = booking_grpc.NewImplementation(s.BookingService(ctx))
	}

	return s.bookgrpcImpl
}

func (s *serviceProvider) HTTPBookingImpl(ctx context.Context) *booking_http.Implementation {
	if s.bookhttpImpl == nil {
		s.bookhttpImpl = booking_http.NewImplementation(s.BookingService(ctx))
	}

	return s.bookhttpImpl
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
	var level zapcore.Level
	if err := level.Set(*logLevel); err != nil {
		log.Fatalf("failed to set log level: %v", err)
	}

	return zap.NewAtomicLevelAt(level)
}
