package booking

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/booking/model"
)

func (s *serv) Create(ctx context.Context, info *model.BookInfo) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.bookRepository.Create(ctx, info)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.bookRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
