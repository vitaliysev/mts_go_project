package hotel

import (
	"context"
	"fmt"
	"github.com/vitaliysev/mts_go_project/internal/hotel/converter"
	desc "github.com/vitaliysev/mts_go_project/pkg/hotel_v1"
)

func (s *serv) GetInfo(ctx context.Context, id int64) (*desc.HotelInfo, error) {
	fmt.Println("bbbbbbbb")
	data, err := s.hotelRepository.GetInfo(ctx, id)
	fmt.Println("ddddd")
	if err != nil {
		return nil, err
	}
	return converter.ToHotelInfoDescFromService(*data), nil
}
