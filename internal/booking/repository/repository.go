package repository

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/booking/model"
)

type BookRepository interface {
	Create(ctx context.Context, info *model.BookInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Book, error)
	List(ctx context.Context, offset, limit, hotel_id int64) ([]*model.Book, error)
	Update(ctx context.Context, id int64, info *model.BookInfo) error
}
