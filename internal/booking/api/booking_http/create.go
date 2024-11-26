package booking_http

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"strings"
	"time"

	hotelv1 "github.com/vitaliysev/mts_go_project/pkg/hotel_v1"
)

func (i *Implementation) Create(ctx context.Context, req *CreateBookingRequest) (*CreateBookingResponse, error) {
	id, err := i.bookService.Create(ctx, req.GetInfo())
	if err != nil {
		return nil, err
	}

	log.Printf("inserted book with id: %d", id)

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("не удалось подключиться: %v", err)
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
	if err1 != nil || err2 != nil {
		log.Fatalf("ошибка парсинга даты: %v", err)
	}
	diff := endDate.Sub(startDate)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetInfo(ctx, req_grpc)
	if err != nil {
		log.Fatalf("ошибка при вызове GetInfo: %v", err)
	}

	return &CreateBookingResponse{
		ID:       id,
		cost:     int64(diff) * resp.GetHotel().GetPrice(),
		title:    resp.GetHotel().GetName(),
		location: resp.GetHotel().GetLocation(),
		period:   int64(diff),
	}, nil
}
