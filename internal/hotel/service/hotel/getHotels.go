package hotel

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/hotel/api/hotel/model"
	"github.com/vitaliysev/mts_go_project/internal/hotel/converter"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
)

func (s *serv) GetHotels(ctx context.Context) ([]model.Hotel, error) {
	ctx, span := tracing.Tracer.Tracer("Hotel-service").Start(ctx, "Service-layer")
	defer span.End()
	hotels, err := s.hotelRepository.GetHotels(ctx)
	ans := make([]model.Hotel, len(hotels))
	if err != nil {
		return nil, err
	}
	for i, hotel := range hotels {
		ans[i] = *converter.ToHotelFromService(&hotel)
	}
	return ans, nil
}
