package hotel

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vitaliysev/mts_go_project/internal/hotel/repository"
	"github.com/vitaliysev/mts_go_project/internal/hotel/repository/hotel/converter"
	modelRepo "github.com/vitaliysev/mts_go_project/internal/hotel/repository/hotel/model"
	"github.com/vitaliysev/mts_go_project/internal/hotel/service/hotel/model"
)

const (
	tableName      = "hotels"
	idColumn       = "id"
	nameColumn     = "hotel_name"
	locationColumn = "location"
	priceColumn    = "price"
	usernameColumn = "username"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.HotelRepository {
	return &repo{db: db}
}

func (r *repo) SaveHotel(ctx context.Context, info *model.HotelInfo, username string) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, locationColumn, priceColumn, usernameColumn).
		Values(info.Name, info.Location, info.Price, username)
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) GetHotels(ctx context.Context) ([]model.Hotel, error) {
	builder := sq.Select(idColumn, nameColumn, locationColumn, priceColumn).
		From(tableName)
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Query(ctx, query, args...)
	ans := make([]model.Hotel, 0)
	for rows.Next() {
		var hotel modelRepo.Hotel
		err = rows.Scan(&hotel.ID, &hotel.Info.Name, &hotel.Info.Location, &hotel.Info.Price)
		if err != nil {
			return nil, err
		}
		ans = append(ans, *converter.ToHotelFromRepo(&hotel))
	}
	return ans, nil
}

func (r *repo) GetHotel(ctx context.Context, id int64) (*model.Hotel, error) {
	builder := sq.Select(idColumn, nameColumn, locationColumn, priceColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id})
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	var hotel modelRepo.Hotel
	err = r.db.QueryRow(ctx, query, args...).Scan(&hotel.ID, &hotel.Info.Name, &hotel.Info.Location, &hotel.Info.Price)
	if err != nil {
		return nil, err
	}
	return converter.ToHotelFromRepo(&hotel), nil
}

func (r *repo) UpdateHotel(ctx context.Context, hotel *model.Hotel) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(nameColumn, hotel.Info.Name).
		Set(locationColumn, hotel.Info.Location).
		Set(priceColumn, hotel.Info.Price).
		Where(sq.Eq{idColumn: hotel.ID})
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) GetInfo(ctx context.Context, id int64) (*model.HotelInfo, error) {
	builder := sq.Select(nameColumn, locationColumn, priceColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	var info modelRepo.HotelInfo
	err = r.db.QueryRow(ctx, query, args...).Scan(&info.Name, &info.Location, &info.Price)
	fmt.Println("ccccc")
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return converter.ToHotelInfoFromRepo(info), nil
}

func (r *repo) GetId(ctx context.Context, username string) (*[]int64, error) {
	builder := sq.Select(idColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{usernameColumn: username}).
		Limit(1000000)
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Создаем срез для хранения всех результатов
	var ans []int64

	// Проходим по всем строкам и заполняем срез
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ans = append(ans, id)
	}
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		return nil, err
	}
	return &ans, nil
}
