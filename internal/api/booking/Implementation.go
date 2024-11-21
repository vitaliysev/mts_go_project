package booking

import (
	"github.com/vitaliysev/mts_go_project/internal/service"
	desc "github.com/vitaliysev/mts_go_project/pkg/booking_v1"
)

type Implementation struct {
	desc.UnimplementedBookingV1Server
	bookService service.BookService
}

func NewImplementation(bookService service.BookService) *Implementation {
	return &Implementation{
		bookService: bookService,
	}
}
