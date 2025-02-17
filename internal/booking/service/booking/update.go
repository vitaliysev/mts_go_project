package booking

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/auth/logger"
	"github.com/vitaliysev/mts_go_project/internal/booking/model"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
)

func (s *serv) Update(ctx context.Context, id int64, info *model.BookInfo) error {
	ctx, span := tracing.Tracer.Tracer("Booking-service").Start(ctx, "Booking.Service.Update")
	defer span.End()
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.bookRepository.Update(ctx, id, info)
		if errTx != nil {
			return errTx
		}

		_, errGet := s.bookRepository.Get(ctx, id)
		if errGet != nil {
			return errTx
		}
		logger.Info("Updated")

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
