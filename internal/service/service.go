package service

import (
	"context"

	"github.com/vitaliysev/mts_go_project/internal/model"
)

type BookService interface {
	Create(ctx context.Context, info *model.BookInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Book, error)
	List(ctx context.Context, offset, limit int64, hotel string) ([]*model.Book, error)
	Update(ctx context.Context, id int64, info *model.BookInfo) error
}
