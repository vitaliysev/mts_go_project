package hotel

import (
	"context"
	desc "github.com/vitaliysev/mts_go_project/pkg/hotel_v1"
)

func (i *Implementation) GetId(ctx context.Context, req *desc.GetIdRequest) (*desc.GetIdResponse, error) {
	data, err := i.HotelService.GetId(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &desc.GetIdResponse{
		Id: *data,
	}, nil
}
