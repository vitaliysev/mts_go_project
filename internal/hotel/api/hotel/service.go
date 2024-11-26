package hotel

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/service"
	desc "github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/pkg/hotel_v1"
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
