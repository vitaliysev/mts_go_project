package service

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/hotel/api/hotel/model"
	desc "github.com/vitaliysev/mts_go_project/pkg/hotel_v1"
)

type HotelService interface {
	SaveHotel(ctx context.Context, info *model.HotelInfo) error
	GetHotels(ctx context.Context) ([]model.Hotel, error)
	GetHotel(ctx context.Context, id int64) (*model.Hotel, error)
	UpdateHotel(ctx context.Context, hotel model.Hotel) error
	GetInfo(ctx context.Context, id int64) (*desc.HotelInfo, error)
}
