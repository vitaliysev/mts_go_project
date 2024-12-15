package auth_grpc

import (
	config2 "github.com/vitaliysev/mts_go_project/internal/auth/config"
	"github.com/vitaliysev/mts_go_project/internal/auth/logger"
	"github.com/vitaliysev/mts_go_project/internal/auth/service"
	descAuth "github.com/vitaliysev/mts_go_project/pkg/auth_v1"
	"go.uber.org/zap"
)

type Implementation struct {
	descAuth.UnimplementedAuthV1Server
	authService service.AuthService

	tokenConfig config2.TokenConfig
}

func NewImplementation(authService service.AuthService) *Implementation {
	token, err := config2.NewTokenConfig()
	if err != nil {
		logger.Error("No config for auth and access", zap.Error(err))
	}
	return &Implementation{
		authService: authService,
		tokenConfig: token,
	}
}
