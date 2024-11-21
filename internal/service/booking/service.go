package booking

import (
	"github.com/vitaliysev/mts_go_project/internal/client/db"
	"github.com/vitaliysev/mts_go_project/internal/repository"
	"github.com/vitaliysev/mts_go_project/internal/service"
)

type serv struct {
	bookRepository repository.BookRepository
	txManager      db.TxManager
}

func NewService(
	bookRepository repository.BookRepository,
	txManager db.TxManager,
) service.BookService {
	return &serv{
		bookRepository: bookRepository,
		txManager:      txManager,
	}
}
