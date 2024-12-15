package booking_http

import (
	"context"
	"fmt"
	"github.com/vitaliysev/mts_go_project/internal/booking/logger"
	descAccess "github.com/vitaliysev/mts_go_project/pkg/access_v1"
	hotelv1 "github.com/vitaliysev/mts_go_project/pkg/hotel_v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
)

func (i *Implementation) Get(ctx context.Context, req *GetBookingRequest, path string) (*GetBookingResponse, error) {
	accessToken := req.Access_token

	ctx_curr := context.Background()
	md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx_curr = metadata.NewOutgoingContext(ctx_curr, md)

	conn, err := grpc.Dial(
		"localhost:50055",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial GRPC client: %v", err)
	}

	cl := descAccess.NewAccessV1Client(conn)
	defer conn.Close()
	username, errReq := cl.Check(ctx_curr, &descAccess.CheckRequest{
		EndpointAddress: path,
	})
	if errReq != nil {
		log.Fatal(errReq.Error())
	}
	logger.Info("Access granted")
	connec, errHot := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if errHot != nil {
		logger.Error("не удалось подключиться: %v", zap.Error(errHot))
	}
	defer connec.Close()

	client := hotelv1.NewHotelV1Client(connec)

	// Создаем запрос
	req_grpc := &hotelv1.GetIdRequest{
		Username: username.GetUsername(),
	}

	id, err := client.GetId(ctx, req_grpc)
	if err != nil {
		logger.Error("ошибка при вызове GetInfo: %v", zap.Error(err))
	}
	var idParam []int64

	if path == "/booking/v1/listCl" {
		idParam = []int64{0}
	} else {
		idParam = id.GetId()
	}
	fmt.Println(idParam)
	ans, err := i.bookService.List(ctx, 0, 10000000, idParam, username.GetUsername())
	if err != nil {
		logger.Error("")
		return nil, err
	}

	logger.Info("Found books %d")

	return &GetBookingResponse{
		Info: ans,
	}, nil
}
