package hotel

import (
	"github.com/vitaliysev/mts_go_project/internal/hotel/repository"
	"github.com/vitaliysev/mts_go_project/internal/hotel/service"
)

type serv struct {
	hotelRepository repository.HotelRepository
}

func NewService(hotelRepository repository.HotelRepository) service.HotelService {
	return &serv{
		hotelRepository: hotelRepository,
	}
}
