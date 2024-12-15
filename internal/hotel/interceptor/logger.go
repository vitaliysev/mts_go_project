package interceptor

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/vitaliysev/mts_go_project/internal/lib/logger"
)

func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	now := time.Now()

	res, err := handler(ctx, req)
	if err != nil {
		logger.Error(err.Error(), zap.String("method", info.FullMethod), zap.Any("req", req))
		return res, err
	}

	logger.Info("requested successfully", zap.String("method", info.FullMethod), zap.Any("req", req), zap.Any("res", res), zap.Duration("duration", time.Since(now)))

	return res, err
}
