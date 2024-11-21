package converter

import (
	"github.com/vitaliysev/mts_go_project/internal/model"
	modelRepo "github.com/vitaliysev/mts_go_project/internal/repository/booking/model"
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
		Title:      info.Title,
		Period_use: info.Period_use,
	}
}
