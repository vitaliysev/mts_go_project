package hotel

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/hotel/converter"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
	desc "github.com/vitaliysev/mts_go_project/pkg/hotel_v1"
)

func (s *serv) GetInfo(ctx context.Context, id int64) (*desc.HotelInfo, error) {
	ctx, span := tracing.Tracer.Tracer("Hotel-service").Start(ctx, "Service.GetInfo")
	defer span.End()
	data, err := s.hotelRepository.GetInfo(ctx, id)
	if err != nil {
		return nil, err
	}
	return converter.ToHotelInfoDescFromService(*data), nil
}
