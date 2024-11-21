package app

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/api/booking"
	"log"

	"github.com/vitaliysev/mts_go_project/internal/client/db"
	"github.com/vitaliysev/mts_go_project/internal/client/db/pg"
	"github.com/vitaliysev/mts_go_project/internal/client/db/transaction"
	"github.com/vitaliysev/mts_go_project/internal/closer"
	"github.com/vitaliysev/mts_go_project/internal/config"
	"github.com/vitaliysev/mts_go_project/internal/repository"
	bookRepository "github.com/vitaliysev/mts_go_project/internal/repository/booking"
	"github.com/vitaliysev/mts_go_project/internal/service"
	bookService "github.com/vitaliysev/mts_go_project/internal/service/booking"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig
	httpConfig config.HTTPConfig

	dbClient       db.Client
	txManager      db.TxManager
	bookRepository repository.BookRepository

	bookService service.BookService

	bookImpl *booking.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
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

func (s *serviceProvider) BookingImpl(ctx context.Context) *booking.Implementation {
	if s.bookImpl == nil {
		s.bookImpl = booking.NewImplementation(s.BookingService(ctx))
	}

	return s.bookImpl
}
