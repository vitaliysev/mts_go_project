package hotel

import (
	"context"
	"fmt"
	desc "github.com/vitaliysev/mts_go_project/pkg/hotel_v1"
)

func (i *Implementation) GetInfo(ctx context.Context, req *desc.GetInfoRequest) (*desc.GetInfoResponse, error) {
	fmt.Println("aaaaaaaa")
	data, err := i.HotelService.GetInfo(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &desc.GetInfoResponse{
		Hotel: data,
	}, nil
}
