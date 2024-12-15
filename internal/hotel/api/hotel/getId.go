package hotel

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/lib/logger"
	desc "github.com/vitaliysev/mts_go_project/pkg/hotel_v1"
	"go.uber.org/zap"
)

func (i *Implementation) GetId(ctx context.Context, req *desc.GetIdRequest) (*desc.GetIdResponse, error) {
	const op = "hotel.handlers.GetId"
	log := logger.With(
		zap.String("op", op),
	)
	log.Info("getting hotel id...")
	data, err := i.HotelService.GetId(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &desc.GetIdResponse{
		Id: *data,
	}, nil
}
