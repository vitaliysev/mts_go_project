package booking

import (
	"github.com/vitaliysev/mts_go_project/internal/booking/client/db"
	"github.com/vitaliysev/mts_go_project/internal/booking/repository"
	"github.com/vitaliysev/mts_go_project/internal/booking/service"
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
