package booking_grpc

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/booking/converter"
	"log"

	desc "github.com/vitaliysev/mts_go_project/pkg/booking_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	bookObj, err := i.bookService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, title: %s, body: %s, created_at: %v, updated_at: %v\n", bookObj.ID, bookObj.Info.Title, bookObj.Info.Period_use, bookObj.CreatedAt, bookObj.UpdatedAt)

	return &desc.GetResponse{
		Book: converter.ToBookFromService(bookObj),
	}, nil
}
