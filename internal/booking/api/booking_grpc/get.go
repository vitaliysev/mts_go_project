package booking_grpc

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/booking/converter"
	desc "github.com/vitaliysev/mts_go_project/pkg/booking_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	bookObj, err := i.bookService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Book: converter.ToBookFromService(bookObj),
	}, nil
}
