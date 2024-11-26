package hotel

import (
	"context"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/api/hotel/model"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/converter"
)

func (s *serv) UpdateHotel(ctx context.Context, hotel model.Hotel) error {
	err := s.hotelRepository.UpdateHotel(ctx, converter.ToHotelServFromApi(&hotel))
	if err != nil {
		return err
	}
	return nil
}
