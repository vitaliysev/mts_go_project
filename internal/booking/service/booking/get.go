package booking

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/booking/model"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.Book, error) {
	ctx, span := tracing.Tracer.Tracer("Booking-service").Start(ctx, "Booking.Service.Get")
	defer span.End()
	book, err := s.bookRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return book, nil
}
