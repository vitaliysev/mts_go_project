package booking

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/vitaliysev/mts_go_project/internal/booking/client/db"
	"github.com/vitaliysev/mts_go_project/internal/booking/model"
	"github.com/vitaliysev/mts_go_project/internal/booking/repository"
	"github.com/vitaliysev/mts_go_project/internal/booking/repository/booking/converter"
	modelRepo "github.com/vitaliysev/mts_go_project/internal/booking/repository/booking/model"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
	"go.opentelemetry.io/otel/codes"
)

const (
	tableName = "booking"

	idColumn        = "id"
	peroidColumn    = "period_use"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
	hotelIdColumn   = "hotel_id"
	usernameColumn  = "username"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.BookRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.BookInfo, username string) (int64, error) {
	ctx, span := tracing.Tracer.Tracer("Booking-service").Start(ctx, "Booking.Repo.Create")
	defer span.End()
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(peroidColumn, hotelIdColumn, usernameColumn).
		Values(info.Period_use, info.Hotel_id, username).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return 0, err
	}

	q := db.Query{
		Name:     "book_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Book, error) {
	ctx, span := tracing.Tracer.Tracer("Booking-service").Start(ctx, "Booking.Repo.Get")
	defer span.End()
	builder := sq.Select(idColumn, peroidColumn, createdAtColumn, updatedAtColumn, hotelIdColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	q := db.Query{
		Name:     "book_repository.Get",
		QueryRaw: query,
	}

	var book modelRepo.Book
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&book.ID, &book.Info.Period_use, &book.CreatedAt, &book.UpdatedAt, &book.Info.Hotel_id)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return converter.ToBookFromRepo(&book), nil
}

func (r *repo) List(ctx context.Context, offset, limit int64, hotel_id []int64, username string) ([]*model.Book, error) {
	ctx, span := tracing.Tracer.Tracer("Booking-service").Start(ctx, "Booking.Repo.List")
	defer span.End()
	var builder sq.SelectBuilder
	if hotel_id[0] != 0 {
		builder = sq.Select(idColumn, peroidColumn, createdAtColumn, updatedAtColumn, hotelIdColumn).
			PlaceholderFormat(sq.Dollar).
			From(tableName).
			Where(sq.Eq{hotelIdColumn: hotel_id}).
			Limit(uint64(limit)).
			Offset(uint64(offset))
	} else {
		builder = sq.Select(idColumn, peroidColumn, createdAtColumn, updatedAtColumn, hotelIdColumn).
			PlaceholderFormat(sq.Dollar).
			From(tableName).
			Where(sq.Eq{usernameColumn: username}).
			Limit(uint64(limit)).
			Offset(uint64(offset))
	}
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
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	defer rows.Close()

	var books []*model.Book
	for rows.Next() {
		var book modelRepo.Book
		if err := rows.Scan(&book.ID, &book.Info.Period_use, &book.CreatedAt, &book.UpdatedAt, &book.Info.Hotel_id); err != nil {
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
		books = append(books, converter.ToBookFromRepo(&book))
	}

	if rows.Err() != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, rows.Err()
	}

	return books, nil
}

func (r *repo) Update(ctx context.Context, id int64, info *model.BookInfo) error {
	ctx, span := tracing.Tracer.Tracer("Booking-service").Start(ctx, "Booking.Repo.Update")
	defer span.End()
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(hotelIdColumn, info.Hotel_id).
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
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		span.SetStatus(codes.Error, err.Error())
		return fmt.Errorf("no rows updated for id: %d", id)
	}

	return nil
}
