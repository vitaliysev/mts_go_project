package auth

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/auth/model"
)

func (s *serv) Get(ctx context.Context, login string) (*model.Auth, error) {
	auth, err := s.authRepository.Get(ctx, login)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
