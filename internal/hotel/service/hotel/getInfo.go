package hotel

import (
	"context"
	"fmt"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/converter"
	desc "github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/pkg/hotel_v1"
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
