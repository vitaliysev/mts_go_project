package model

import (
	"database/sql"
	"time"
)

type Book struct {
	ID        int64        `json:"id"`
	Info      BookInfo     `json:"info"`
	CreatedAt time.Time    `json:"create_time"`
	UpdatedAt sql.NullTime `json:"update_time"`
}

type BookInfo struct {
	Period_use string `json:"period_use"`
	Hotel_id   int64  `json:"hotel_id"`
}
