package models

type CreateBookingResponse struct {
	ID       int64  `json:"id"`
	Cost     int64  `json:"cost"`
	Title    string `json:"title"`
	Location string `json:"location"`
	Period   int64  `json:"period"`
}
