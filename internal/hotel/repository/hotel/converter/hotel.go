package converter

import (
	modelRepo "github.com/vitaliysev/mts_go_project/internal/hotel/repository/hotel/model"
	"github.com/vitaliysev/mts_go_project/internal/hotel/service/hotel/model"
)

func ToHotelFromRepo(hotel *modelRepo.Hotel) *model.Hotel {
	return &model.Hotel{
		ID:   hotel.ID,
		Info: *ToHotelInfoFromRepo(hotel.Info),
	}
}

func ToHotelInfoFromRepo(info modelRepo.HotelInfo) *model.HotelInfo {
	return &model.HotelInfo{
		Name:     info.Name,
		Location: info.Location,
		Price:    info.Price,
	}
}
