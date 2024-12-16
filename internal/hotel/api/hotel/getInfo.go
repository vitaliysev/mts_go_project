package hotel

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/lib/logger"
	desc "github.com/vitaliysev/mts_go_project/pkg/hotel_v1"
	"go.uber.org/zap"
)

func (i *Implementation) GetInfo(ctx context.Context, req *desc.GetInfoRequest) (*desc.GetInfoResponse, error) {
	const op = "hotel.GetInfo"
	log := logger.With(
		zap.String("op", op),
	)
	log.Info("getting hotel info...")
	data, err := i.HotelService.GetInfo(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &desc.GetInfoResponse{
		Hotel: data,
	}, nil
}
