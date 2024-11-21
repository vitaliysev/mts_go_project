package booking

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/model"
)

func (s *serv) List(ctx context.Context, offset, limit int64, hotel string) ([]*model.Book, error) {
	var books []*model.Book

	// Выполняем операцию в транзакции (если нужно)
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		books, errTx = s.bookRepository.List(ctx, offset, limit, hotel)
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
