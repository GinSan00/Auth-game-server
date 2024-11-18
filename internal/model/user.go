package model

import "time"

type User struct {
	ID        string
	Info      UserInfo
	CreatedAt time.Time
}

type UserInfo struct {
	Email    string
	Password string
	Nickname string
}
