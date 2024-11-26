package hotel

import (
	"context"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/api/hotel/model"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/converter"
)

func (s *serv) GetHotel(ctx context.Context, id int64) (*model.Hotel, error) {
	hotel, err := s.hotelRepository.GetHotel(ctx, id)
	if err != nil {
		return nil, err
	}
	return converter.ToHotelFromService(hotel), nil
}
