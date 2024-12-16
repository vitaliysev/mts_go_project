package booking_http

import (
	"context"
	"fmt"
	"github.com/vitaliysev/mts_go_project/internal/booking/logger"
	metric "github.com/vitaliysev/mts_go_project/internal/booking/metrics"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"time"

	authv1 "github.com/vitaliysev/mts_go_project/pkg/auth_v1"
)

func (i *Implementation) Login(ctx context.Context, req *LoginClientRequest) (*LoginClientResponse, error) {
	timeStart := time.Now()
	metric.IncRequestCounter()
	ctx, span := tracing.Tracer.Tracer("Booking-service").Start(ctx, "booking.Login")
	defer span.End()
	traceId := fmt.Sprintf("%s", span.SpanContext().TraceID())

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
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceId)

	resp, err := client.Login(ctx, req_grpc)
	if err != nil {
		diffTime := time.Since(timeStart)
		metric.IncResponseCounter("error", "booking.Login")
		metric.HistogramResponseTimeObserve("error", diffTime.Seconds())
		logger.Error("ошибка при вызове", zap.Error(err))
	}
	diffTime := time.Since(timeStart)
	metric.HistogramResponseTimeObserve("success", diffTime.Seconds())
	metric.IncResponseCounter("success", "booking.Login")
	return &LoginClientResponse{
		Access_token: resp.AccessToken,
	}, nil
}
