package booking_http

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/booking/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"

	authv1 "github.com/vitaliysev/mts_go_project/pkg/auth_v1"
)

func (i *Implementation) Login(ctx context.Context, req *LoginClientRequest) (*LoginClientResponse, error) {
	conn, err := grpc.Dial("localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("не удалось подключиться: %v", zap.Error(err))
	}
	defer conn.Close()

	client := authv1.NewAuthV1Client(conn)

	// Создаем запрос
	req_grpc := &authv1.LoginRequest{
		Username:     req.Username,
		Password:     req.Password,
		RefreshToken: req.Refresh_token,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Login(ctx, req_grpc)
	if err != nil {
		logger.Error("ошибка при вызове", zap.Error(err))
	}

	return &LoginClientResponse{
		Access_token: resp.AccessToken,
	}, nil
}
