package booking

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/booking/model"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
)

func (s *serv) List(ctx context.Context, offset, limit int64, hotel_id []int64, username string) ([]*model.Book, error) {
	ctx, span := tracing.Tracer.Tracer("Booking-service").Start(ctx, "Booking.List.Service-layer")
	defer span.End()
	var books []*model.Book

	// Выполняем операцию в транзакции (если нужно)
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		books, errTx = s.bookRepository.List(ctx, offset, limit, hotel_id, username)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return books, nil
}
