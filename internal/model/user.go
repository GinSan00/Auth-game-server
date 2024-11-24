package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
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
