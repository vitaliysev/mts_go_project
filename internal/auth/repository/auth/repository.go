package auth

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/vitaliysev/mts_go_project/internal/auth/client/db"
	modelRepo "github.com/vitaliysev/mts_go_project/internal/auth/model"
	"github.com/vitaliysev/mts_go_project/internal/auth/repository"
	"github.com/vitaliysev/mts_go_project/internal/auth/repository/auth/converter"
	"github.com/vitaliysev/mts_go_project/internal/tracing"
)

const (
	tableName = "auth"

	loginColumn    = "login"
	passwordColumn = "hashed_password"
	roleColumn     = "role"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.AuthRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *modelRepo.AuthInfo) (string, error) {
	ctx, span := tracing.Tracer.Tracer("Auth-service").Start(ctx, "Auth.Repo.Create")
	defer span.End()
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(loginColumn, passwordColumn, roleColumn).
		Values(info.Login, info.Hashed_password, info.Role).
		Suffix("RETURNING login")

	query, args, err := builder.ToSql()
	if err != nil {
		return "nothing", err
	}

	q := db.Query{
		Name:     "auth_repository.Create",
		QueryRaw: query,
	}

	var login string
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&login)
	if err != nil {
		return "nothing", fmt.Errorf("failed to execute query: %w", err)
	}

	return login, nil
}

func (r *repo) Get(ctx context.Context, login string) (*modelRepo.Auth, error) {
	ctx, span := tracing.Tracer.Tracer("Auth-service").Start(ctx, "Auth.Repo.Get")
	defer span.End()
	builder := sq.Select(loginColumn, passwordColumn, roleColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{loginColumn: login}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "auth_repository.Get",
		QueryRaw: query,
	}

	var auth modelRepo.Auth
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&auth.Info.Login, &auth.Info.Hashed_password, &auth.Info.Role)
	if err != nil {
		return nil, err
	}

	return converter.ToAuthFromRepo(&auth), nil
}
