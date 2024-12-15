package auth_grpc

import (
	"context"
	"fmt"
	"github.com/vitaliysev/mts_go_project/internal/auth/logger"
	"github.com/vitaliysev/mts_go_project/internal/auth/model"
	"github.com/vitaliysev/mts_go_project/internal/auth/utils"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
	descAuth "github.com/vitaliysev/mts_go_project/pkg/auth_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Login(ctx context.Context, req *descAuth.LoginRequest) (*descAuth.LoginResponse, error) {
	const op = "Auth.Login"
	ctx, span := tracing.Tracer.Tracer("Auth-service").Start(ctx, op)
	defer span.End()
	traceId := fmt.Sprintf("%s", span.SpanContext().TraceID())
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceId)
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
