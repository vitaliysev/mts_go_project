package booking

import (
	"context"
	"log"

	"github.com/vitaliysev/mts_go_project/internal/converter"
	desc "github.com/vitaliysev/mts_go_project/pkg/booking_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.bookService.Create(ctx, converter.ToBookInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted book with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
