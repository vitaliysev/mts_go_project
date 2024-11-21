package booking

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"

	"github.com/vitaliysev/mts_go_project/internal/converter"
	desc "github.com/vitaliysev/mts_go_project/pkg/booking_v1"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	// Преобразование данных из запроса в формат, используемый сервисом
	updateInfo := converter.ToUpdateBookInfoFromDesc(req.GetInfo())

	// Обновление информации о книге в сервисе
	err := i.bookService.Update(ctx, req.GetId(), updateInfo)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	log.Printf("updated book with id: %d", req.GetId())

	return &emptypb.Empty{}, nil
}
