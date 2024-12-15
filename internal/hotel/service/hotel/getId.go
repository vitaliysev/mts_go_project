package hotel

import (
	"context"
	"fmt"
)

func (s *serv) GetId(ctx context.Context, username string) (*[]int64, error) {
	fmt.Println("bbbbbbbb")
	data, err := s.hotelRepository.GetId(ctx, username)
	if err != nil {
		return nil, err
	}
	return data, nil
}
