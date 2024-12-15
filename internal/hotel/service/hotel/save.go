package hotel

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/hotel/api/hotel/model"
	"github.com/vitaliysev/mts_go_project/internal/hotel/converter"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
)

func (s *serv) SaveHotel(ctx context.Context, info *model.HotelInfo, username string) error {
	ctx, span := tracing.Tracer.Tracer("Hotel-service").Start(ctx, "Service-layer")
	defer span.End()
	err := s.hotelRepository.SaveHotel(ctx, converter.ToHotelServInfoFromApi(info), username)
	if err != nil {
		return err
	}
	return nil
}
