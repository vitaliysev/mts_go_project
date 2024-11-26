package app

import (
	"context"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/api/hotel"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/closer"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/config"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/repository"
	hotelRepository "github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/repository/hotel"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/service"
	hotelService "github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/service/hotel"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig
	restConfig config.RESTConfig

	pgPool          *pgxpool.Pool
	hotelRepository repository.HotelRepository
	hotelService    service.HotelService

	hotelImpl *hotel.Implementation
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

func (s *serviceProvider) RESTConfig() config.RESTConfig {
	if s.restConfig == nil {
		cfg, err := config.NewRESTConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}
		s.restConfig = cfg
	}
	return s.restConfig
}

func (s *serviceProvider) PgPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database")
		}
		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping database")
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})
		s.pgPool = pool
	}
	return s.pgPool

}

func (s *serviceProvider) HotelRepository(ctx context.Context) repository.HotelRepository {
	if s.hotelRepository == nil {
		s.hotelRepository = hotelRepository.NewRepository(s.PgPool(ctx))
	}
	return s.hotelRepository
}

func (s *serviceProvider) HotelService(ctx context.Context) service.HotelService {
	if s.hotelService == nil {
		s.hotelService = hotelService.NewService(s.HotelRepository(ctx))
	}
	return s.hotelService
}

func (s *serviceProvider) HotelImpl(ctx context.Context) *hotel.Implementation {
	if s.hotelImpl == nil {
		s.hotelImpl = hotel.NewImplementation(s.HotelService(ctx))
	}
	return s.hotelImpl
}
