package booking

import (
	"context"

	"github.com/vitaliysev/mts_go_project/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.Book, error) {
	book, err := s.bookRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return book, nil
}
