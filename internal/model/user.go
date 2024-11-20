package model

import (
	"time"
)

type User struct {
	ID        string
	Info      UserInfo
	CreatedAt time.Time
	LastLogin time.Time
}

type UserInfo struct {
	Email    string
	PassHash []byte
	Nickname string
	Elo      uint
}
