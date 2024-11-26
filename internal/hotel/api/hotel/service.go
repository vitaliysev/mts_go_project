package hotel

import (
	"github.com/vitaliysev/mts_go_project/internal/hotel/service"
	desc "github.com/vitaliysev/mts_go_project/pkg/hotel_v1"
)

type Implementation struct {
	desc.UnimplementedHotelV1Server
	service.HotelService
}

func NewImplementation(hotelService service.HotelService) *Implementation {
	return &Implementation{
		HotelService: hotelService,
	}
}
