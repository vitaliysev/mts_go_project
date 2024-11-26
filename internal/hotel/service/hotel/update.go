package hotel

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/hotel/api/hotel/model"
	"github.com/vitaliysev/mts_go_project/internal/hotel/converter"
)

func (s *serv) UpdateHotel(ctx context.Context, hotel model.Hotel) error {
	err := s.hotelRepository.UpdateHotel(ctx, converter.ToHotelServFromApi(&hotel))
	if err != nil {
		return err
	}
	return nil
}
