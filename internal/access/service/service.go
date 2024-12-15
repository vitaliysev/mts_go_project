package service

import (
	"context"
)

type AccessService interface {
	AccessibleRoles(ctx context.Context) (map[string]string, error)
}
