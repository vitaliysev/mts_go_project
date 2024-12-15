package auth_grpc

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/auth/logger"
	"github.com/vitaliysev/mts_go_project/internal/auth/model"
	"github.com/vitaliysev/mts_go_project/internal/auth/utils"
	descAuth "github.com/vitaliysev/mts_go_project/pkg/auth_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetRefreshToken(ctx context.Context, req *descAuth.GetRefreshTokenRequest) (*descAuth.GetRefreshTokenResponse, error) {
	claims, err := utils.VerifyToken(req.GetRefreshToken(), []byte(i.tokenConfig.GetRefr()))
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "invalid refresh token")
	}

	currAuth, err := i.authService.Get(ctx, claims.Username)
	if err != nil {
		return nil, err
	}

	logger.Info("Login found")
	refreshToken, err := utils.GenerateToken(model.UserInfo{
		Username: claims.Username,
		Role:     currAuth.Info.Role,
	},
		[]byte(i.tokenConfig.GetRefr()),
		i.tokenConfig.GetRefreshTime(),
	)
	if err != nil {
		return nil, err
	}

	return &descAuth.GetRefreshTokenResponse{RefreshToken: refreshToken}, nil
}
