package hotel

import (
	"context"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/api/hotel/model"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/converter"
)

func (s *serv) SaveHotel(ctx context.Context, info *model.HotelInfo) error {
	err := s.hotelRepository.SaveHotel(ctx, converter.ToHotelServInfoFromApi(info))
	if err != nil {
		return err
	}
	return nil
}
