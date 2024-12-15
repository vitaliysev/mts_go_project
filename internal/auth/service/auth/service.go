package auth

import (
	"github.com/vitaliysev/mts_go_project/internal/auth/client/db"
	"github.com/vitaliysev/mts_go_project/internal/auth/repository"
	"github.com/vitaliysev/mts_go_project/internal/auth/service"
)

type serv struct {
	authRepository repository.AuthRepository
	txManager      db.TxManager
}

func NewService(
	authRepository repository.AuthRepository,
	txManager db.TxManager,
) service.AuthService {
	return &serv{
		authRepository: authRepository,
		txManager:      txManager,
	}
}
