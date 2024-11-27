package booking_http

import (
	"context"
	"log"
)

func (i *Implementation) Get(ctx context.Context, req *GetBookingRequest) (*GetBookingResponse, error) {
	ans, err := i.bookService.List(ctx, 0, 10000000, req.GetInfo())
	if err != nil {
		return nil, err
	}

	log.Printf("Found books %d")

	return &GetBookingResponse{
		Info: ans,
	}, nil
}
