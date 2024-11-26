package model

type HotelInfo struct {
	Name     string
	Location string
	Price    int
}

type Hotel struct {
	ID   int64
	Info HotelInfo
}
