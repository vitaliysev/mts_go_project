package repository

import (
	"context"
	"github.com/vitaliysev/mts_go_project/internal/auth/model"
)

type AuthRepository interface {
	Create(ctx context.Context, info *model.AuthInfo) (string, error)
	Get(ctx context.Context, login string) (*model.Auth, error)
}
