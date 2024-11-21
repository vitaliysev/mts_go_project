package booking

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/model"
)

func (s *serv) Update(ctx context.Context, id int64, info *model.BookInfo) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.bookRepository.Update(ctx, id, info)
		if errTx != nil {
			return errTx
		}

		errTx = s.bookRepository.Update(ctx, id, info)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
