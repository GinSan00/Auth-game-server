package service

import (
	"context"
	"main/internal/model"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, info *model.UserInfo) (string, error)
}
