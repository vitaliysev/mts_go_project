package app

import (
	"context"
	"flag"
	"github.com/natefinch/lumberjack"
	"github.com/vitaliysev/mts_go_project/internal/access/api/access_grpc"
	config2 "github.com/vitaliysev/mts_go_project/internal/access/config"
	"github.com/vitaliysev/mts_go_project/internal/access/service"
	accessService "github.com/vitaliysev/mts_go_project/internal/access/service/access"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var logLevel = flag.String("l", "info", "log level")

type serviceProvider struct {
	pgConfig   config2.PGConfig
	grpcConfig config2.GRPCConfig

	accessService  service.AccessService
	accessgrpcImpl *access_grpc.Implementation
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

func (s *serviceProvider) GRPCAccessConfig() config2.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config2.NewGRPCAccessConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) AccessService(ctx context.Context) service.AccessService {
	if s.accessService == nil {
		s.accessService = accessService.NewService()
	}

	return s.accessService
}

func (s *serviceProvider) GRPCAccessImpl(ctx context.Context) *access_grpc.Implementation {
	if s.accessgrpcImpl == nil {
		s.accessgrpcImpl = access_grpc.NewImplementation(s.AccessService(ctx))
	}

	return s.accessgrpcImpl
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
