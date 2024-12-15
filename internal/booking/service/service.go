package service

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/booking/model"
)

type BookService interface {
	Create(ctx context.Context, info *model.BookInfo, username string) (int64, error)
	Get(ctx context.Context, id int64) (*model.Book, error)
	List(ctx context.Context, offset, limit int64, id []int64, username string) ([]*model.Book, error)
	Update(ctx context.Context, id int64, info *model.BookInfo) error
}
