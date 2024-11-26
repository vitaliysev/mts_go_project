package converter

import (
	modelRepo "github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/repository/hotel/model"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/service/hotel/model"
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
