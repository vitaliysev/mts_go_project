package model

import (
	"database/sql"
	"time"
)

type Book struct {
	ID        int64
	Info      BookInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type BookInfo struct {
	Period_use string
	Hotel_id   int64
}
