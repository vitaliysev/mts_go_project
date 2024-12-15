package hotel

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/hotel/api/hotel/model"
	"github.com/vitaliysev/mts_go_project/internal/hotel/converter"
)

func (s *serv) SaveHotel(ctx context.Context, info *model.HotelInfo, username string) error {
	err := s.hotelRepository.SaveHotel(ctx, converter.ToHotelServInfoFromApi(info), username)
	if err != nil {
		return err
	}
	return nil
}
