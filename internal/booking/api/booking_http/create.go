package booking_http

import (
	"context"
	"fmt"
	"github.com/vitaliysev/mts_go_project/internal/booking/logger"
	metric "github.com/vitaliysev/mts_go_project/internal/booking/metrics"
	"github.com/vitaliysev/mts_go_project/internal/models"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
	descAccess "github.com/vitaliysev/mts_go_project/pkg/access_v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
	"time"

	hotelv1 "github.com/vitaliysev/mts_go_project/pkg/hotel_v1"
)

func (i *Implementation) Create(ctx context.Context, req *CreateBookingRequest) (*models.CreateBookingResponse, error) {
	ctx, span := tracing.Tracer.Tracer("Booking-service").Start(ctx, "booking.Create")
	defer span.End()
	traceId := fmt.Sprintf("%s", span.SpanContext().TraceID())
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", traceId)

	timeStart := time.Now()
	metric.IncRequestCounter()
	accessToken := req.Access_token
	ctx_curr := context.Background()
	md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx_curr = metadata.NewOutgoingContext(ctx_curr, md)
	ctx_curr = metadata.AppendToOutgoingContext(ctx_curr, "x-trace-id", traceId)
	conn, err := grpc.Dial(
		"localhost:50055",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial GRPC client: %v", err)
	}

	cl := descAccess.NewAccessV1Client(conn)
	fmt.Println("aaa")
	username, errReq := cl.Check(ctx_curr, &descAccess.CheckRequest{
		EndpointAddress: "/booking/v1/create",
	})
	if errReq != nil {
		log.Fatal(err.Error())
	}

	logger.Info("Access granted")

	id, err := i.bookService.Create(ctx, req.GetInfo(), username.GetUsername())
	if err != nil {
		return nil, err
	}
	conn.Close()

	conn, err = grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		logger.Error("не удалось подключиться: %v", zap.Error(err))
	}
	defer conn.Close()

	client := hotelv1.NewHotelV1Client(conn)

	// Создаем запрос
	req_grpc := &hotelv1.GetInfoRequest{
		Id: req.GetInfo().Hotel_id,
	}
	dates_range := req.GetInfo().Period_use
	dates := strings.Split(dates_range, "-")
	start := dates[0]
	end := dates[1]
	layout := "02.01.2006"
	startDate, err1 := time.Parse(layout, start)
	endDate, err2 := time.Parse(layout, end)
	if err1 != nil {
		logger.Error("ошибка парсинга даты: %v", zap.Error(err1))
	}
	if err2 != nil {
		logger.Error("ошибка парсинга даты: %v", zap.Error(err2))
	}
	if startDate.After(endDate) {
		logger.Error("дата заезда позднее даты выезда")
	}
	diff := endDate.Sub(startDate)/(1000000000*3600*24) + 1

	resp, err := client.GetInfo(ctx, req_grpc)
	if err != nil {
		logger.Error("ошибка при вызове GetInfo: %v", zap.Error(err))
	}
	diffTime := time.Since(timeStart)
	metric.IncResponseCounter("success", "booking.Get")
	metric.HistogramResponseTimeObserve("success", diffTime.Seconds())
	logger.Info("inserted book with %d", zap.Int64("id", id))
	response := models.CreateBookingResponse{
		ID:       id,
		Cost:     int64(diff) * resp.GetHotel().GetPrice(),
		Title:    resp.GetHotel().GetName(),
		Location: resp.GetHotel().GetLocation(),
		Period:   int64(diff),
	}
	metric.IncMessagesCounter()
	i.producer.SendMessage(response)

	return &response, nil
}
