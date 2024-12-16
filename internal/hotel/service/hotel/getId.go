package hotel

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
)

func (s *serv) GetId(ctx context.Context, username string) (*[]int64, error) {
	ctx, span := tracing.Tracer.Tracer("Hotel-service").Start(ctx, "Service.getId")
	defer span.End()
	data, err := s.hotelRepository.GetId(ctx, username)
	if err != nil {
		return nil, err
	}
	return data, nil
}
