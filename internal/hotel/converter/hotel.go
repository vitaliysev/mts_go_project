package converter

import (
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/api/hotel/model"
	modelServ "github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/service/hotel/model"
	desc "github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/pkg/hotel_v1"
)

func ToHotelFromService(note *modelServ.Hotel) *model.Hotel {
	return &model.Hotel{
		ID:   note.ID,
		Info: *ToHotelInfoFromService(note.Info),
	}
}

func ToHotelInfoFromService(info modelServ.HotelInfo) *model.HotelInfo {
	return &model.HotelInfo{
		Name:     info.Name,
		Location: info.Location,
		Price:    info.Price,
	}
}
func ToHotelInfoDescFromService(info modelServ.HotelInfo) *desc.HotelInfo {
	return &desc.HotelInfo{
		Name:     info.Name,
		Location: info.Location,
		Price:    int64(info.Price),
	}
}
func ToHotelServInfoFromApi(info *model.HotelInfo) *modelServ.HotelInfo {
	return &modelServ.HotelInfo{
		Name:     info.Name,
		Location: info.Location,
		Price:    info.Price,
	}
}
func ToHotelServFromApi(hotel *model.Hotel) *modelServ.Hotel {
	return &modelServ.Hotel{
		ID:   hotel.ID,
		Info: *ToHotelServInfoFromApi(&hotel.Info),
	}
}
