package hotel

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/repository"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/service"
)

type serv struct {
	hotelRepository repository.HotelRepository
}

func NewService(hotelRepository repository.HotelRepository) service.HotelService {
	return &serv{
		hotelRepository: hotelRepository,
	}
}
