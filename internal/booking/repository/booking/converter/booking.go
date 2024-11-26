package converter

import (
	"github.com/vitaliysev/mts_go_project/internal/booking/model"
	modelRepo "github.com/vitaliysev/mts_go_project/internal/booking/repository/booking/model"
)

func ToBookFromRepo(book *modelRepo.Book) *model.Book {
	return &model.Book{
		ID:        book.ID,
		Info:      ToBookInfoFromRepo(book.Info),
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
}

func ToBookInfoFromRepo(info modelRepo.BookInfo) model.BookInfo {
	return model.BookInfo{
		Period_use: info.Period_use,
		Hotel_id:   info.Hotel_id,
	}
}
