package booking

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"

	"github.com/vitaliysev/mts_go_project/internal/client/db"
	"github.com/vitaliysev/mts_go_project/internal/model"
	"github.com/vitaliysev/mts_go_project/internal/repository"
	"github.com/vitaliysev/mts_go_project/internal/repository/booking/converter"
	modelRepo "github.com/vitaliysev/mts_go_project/internal/repository/booking/model"
)

const (
	tableName = "booking"

	idColumn        = "id"
	titleColumn     = "title"
	peroidColumn    = "period_use"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.BookRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.BookInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(titleColumn, peroidColumn).
		Values(info.Title, info.Period_use).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "book_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Book, error) {
	builder := sq.Select(idColumn, titleColumn, peroidColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "book_repository.Get",
		QueryRaw: query,
	}

	var book modelRepo.Book
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&book.ID, &book.Info.Title, &book.Info.Period_use, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToBookFromRepo(&book), nil
}

func (r *repo) List(ctx context.Context, offset, limit int64, hotel string) ([]*model.Book, error) {
	builder := sq.Select(idColumn, titleColumn, peroidColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{titleColumn: hotel}).
		Limit(uint64(limit)).
		Offset(uint64(offset))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "book_repository.List",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*model.Book
	for rows.Next() {
		var book modelRepo.Book
		if err := rows.Scan(&book.ID, &book.Info.Title, &book.Info.Period_use, &book.CreatedAt, &book.UpdatedAt); err != nil {
			return nil, err
		}
		books = append(books, converter.ToBookFromRepo(&book))
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return books, nil
}

func (r *repo) Update(ctx context.Context, id int64, info *model.BookInfo) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(titleColumn, info.Title).
		Set(peroidColumn, info.Period_use).
		Set(updatedAtColumn, sq.Expr("NOW()")). // Обновляем поле `updated_at` текущей датой
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "book_repository.Update",
		QueryRaw: query,
	}

	result, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("no rows updated for id: %d", id)
	}

	return nil
}
