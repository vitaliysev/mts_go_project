package hotel

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/hotel/api/hotel/model"
	"github.com/vitaliysev/mts_go_project/internal/hotel/converter"
)

func (s *serv) GetHotel(ctx context.Context, id int64) (*model.Hotel, error) {
	hotel, err := s.hotelRepository.GetHotel(ctx, id)
	if err != nil {
		return nil, err
	}
	return converter.ToHotelFromService(hotel), nil
}
