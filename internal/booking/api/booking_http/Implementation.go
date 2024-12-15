package booking_http

import (
	"github.com/vitaliysev/mts_go_project/internal/booking/model"
	"github.com/vitaliysev/mts_go_project/internal/booking/redpanda/producer"
	"github.com/vitaliysev/mts_go_project/internal/booking/service"
)

type Implementation struct {
	bookService service.BookService
	producer    producer.Producer
}

func NewImplementation(bookService service.BookService, producer producer.Producer) *Implementation {
	return &Implementation{
		bookService: bookService,
		producer:    producer,
	}
}

type CreateBookingRequest struct {
	Info         model.BookInfo `json:"info"`
	Access_token string         `json:"access_token"`
}

type CreateBookingResponse struct {
	ID       int64  `json:"id"`
	Cost     int64  `json:"cost"`
	Title    string `json:"title"`
	Location string `json:"location"`
	Period   int64  `json:"period"`
}

type SigninClientRequest struct {
	Username string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type SigninClientResponse struct {
	Refresh_token string `json:"refresh_token"`
}

type LoginClientRequest struct {
	Username      string `json:"login"`
	Password      string `json:"password"`
	Refresh_token string `json:"refresh_token"`
}

type LoginClientResponse struct {
	Access_token string `json:"access_token"`
}

// @description GetBookingRequest contains a info for get
type GetBookingRequest struct {
	ID           int64  `json:"id"`
	Access_token string `json:"access_token"`
	Path         string
}

// @description GetBookingResponse contains a list of booking information.
type GetBookingResponse struct {
	Info []*model.Book `json:"info"`
}

type GetRefreshTokenRequest struct {
	Refresh_token string `json:"refresh_token"`
}

type GetRefreshTokenResponse struct {
	Refresh_token string `json:"refresh_token"`
}

func (x *CreateBookingRequest) GetInfo() *model.BookInfo {
	if x != nil {
		return &x.Info
	}
	return nil
}
