package booking_grpc

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/booking/converter"
	desc "github.com/vitaliysev/mts_go_project/pkg/booking_v1"
	"log"
)

func (i *Implementation) List(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
	// Получение списка книг из сервиса
	books, err := i.bookService.List(ctx, req.GetOffset(), req.GetLimit(), req.GetHotelId())
	if err != nil {
		return nil, err
	}

	// Преобразование списка книг в формат ответа
	bookList := make([]*desc.Book, len(books))

	for idx, book := range books {
		bookList[idx] = converter.ToBookFromService(book)
	}

	log.Printf("retrieved %d books", len(bookList))

	return &desc.ListResponse{
		Books: bookList,
	}, nil
}
