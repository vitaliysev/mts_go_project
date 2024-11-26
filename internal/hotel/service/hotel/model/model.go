package model

type Hotel struct {
	ID   int64
	Info HotelInfo
}

type HotelInfo struct {
	Name     string
	Location string
	Price    int
}
