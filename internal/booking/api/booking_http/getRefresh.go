package booking_http

import (
	"context"
	"fmt"
	"github.com/vitaliysev/mts_go_project/internal/booking/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"

	authv1 "github.com/vitaliysev/mts_go_project/pkg/auth_v1"
)

func (i *Implementation) GetRefresh(ctx context.Context, req *GetRefreshTokenRequest) (*GetRefreshTokenResponse, error) {
	conn, err := grpc.Dial("localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("не удалось подключиться: %v", zap.Error(err))
	}
	logger.Info("Connected")
	defer conn.Close()

	client := authv1.NewAuthV1Client(conn)

	// Создаем запрос
	req_grpc := &authv1.GetRefreshTokenRequest{
		RefreshToken: req.Refresh_token,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetRefreshToken(ctx, req_grpc)
	if err != nil {
		logger.Error("ошибка при вызове:", zap.Error(err))
	}

	fmt.Println(resp.RefreshToken)
	return &GetRefreshTokenResponse{
		Refresh_token: resp.RefreshToken,
	}, nil
}
