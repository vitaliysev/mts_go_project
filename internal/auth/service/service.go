package service

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/auth/model"
)

type AuthService interface {
	Create(ctx context.Context, info *model.AuthInfo) (string, error)
	Get(ctx context.Context, login string) (*model.Auth, error)
}
