package model

type Hotel struct {
	ID   int64 `json:"id" validate:"required"`
	Info HotelInfo
}

type HotelInfo struct {
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
	Price    int    `json:"price" validate:"required"`
}
