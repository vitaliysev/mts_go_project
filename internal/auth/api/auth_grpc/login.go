package auth_grpc

import (
	"context"
	"fmt"
	"github.com/vitaliysev/mts_go_project/internal/auth/logger"
	"github.com/vitaliysev/mts_go_project/internal/auth/model"
	"github.com/vitaliysev/mts_go_project/internal/auth/utils"
	descAuth "github.com/vitaliysev/mts_go_project/pkg/auth_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Login(ctx context.Context, req *descAuth.LoginRequest) (*descAuth.LoginResponse, error) {
	claims, err := utils.VerifyToken(req.GetRefreshToken(), []byte(i.tokenConfig.GetRefr()))
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "invalid refresh token")
	}

	currAuth, err := i.authService.Get(ctx, req.GetUsername())
	if err != nil {
		return nil, err
	}

	fmt.Println("Login found")

	if !utils.VerifyPassword(currAuth.Info.Hashed_password, req.GetPassword()) {
		logger.Error("Login failed, invalid password")
	}

	accessToken, err := utils.GenerateToken(model.UserInfo{
		Username: claims.Username,
		Role:     currAuth.Info.Role,
	},
		[]byte(i.tokenConfig.GetAccess()),
		i.tokenConfig.GetAccessTime(),
	)
	if err != nil {
		return nil, err
	}
	fmt.Println(accessToken)
	return &descAuth.LoginResponse{AccessToken: accessToken}, nil
}
