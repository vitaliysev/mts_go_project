package auth

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/auth/model"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
)

func (s *serv) Get(ctx context.Context, login string) (*model.Auth, error) {
	ctx, span := tracing.Tracer.Tracer("Auth-service").Start(ctx, "Auth.Service.Get")
	defer span.End()
	auth, err := s.authRepository.Get(ctx, login)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
