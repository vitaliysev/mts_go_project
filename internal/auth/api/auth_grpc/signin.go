package auth_grpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/vitaliysev/mts_go_project/internal/auth/converter"
	"github.com/vitaliysev/mts_go_project/internal/auth/model"
	"github.com/vitaliysev/mts_go_project/internal/auth/utils"
	descAuth "github.com/vitaliysev/mts_go_project/pkg/auth_v1"
)

func (i *Implementation) Signin(ctx context.Context, req *descAuth.SigninRequest) (*descAuth.SigninResponse, error) {
	i.authService.Create(ctx, converter.ToCreateFromSignin(req))
	fmt.Println("Created")
	refreshToken, err := utils.GenerateToken(model.UserInfo{
		Username: req.GetUsername(),
		Role:     req.GetRole(),
	},
		[]byte(i.tokenConfig.GetRefr()),
		i.tokenConfig.GetRefreshTime(),
	)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}
	fmt.Println(refreshToken)
	return &descAuth.SigninResponse{RefreshToken: refreshToken}, nil
}
