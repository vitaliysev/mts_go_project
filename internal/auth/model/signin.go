package model

import (
	"time"
)

type Auth struct {
	Info      AuthInfo  `json:"info"`
	CreatedAt time.Time `json:"create_time"`
}

type AuthInfo struct {
	Login           string `json:"login"`
	Hashed_password string `json:"password"`
	Role            string `json:"role"`
}
