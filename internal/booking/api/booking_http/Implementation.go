package booking_http

import (
	"github.com/vitaliysev/mts_go_project/internal/booking/model"
	"github.com/vitaliysev/mts_go_project/internal/booking/service"
)

type Implementation struct {
	bookService service.BookService
}

func NewImplementation(bookService service.BookService) *Implementation {
	return &Implementation{
		bookService: bookService,
	}
}

type CreateBookingRequest struct {
	Info model.BookInfo `json:"info"`
}

type CreateBookingResponse struct {
	ID       int64  `json:"id"`
	Cost     int64  `json:"cost"`
	Title    string `json:"title"`
	Location string `json:"location"`
	Period   int64  `json:"period"`
}

func (x *CreateBookingRequest) GetInfo() *model.BookInfo {
	if x != nil {
		return &x.Info
	}
	return nil
}
