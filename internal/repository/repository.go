package repository

import (
	"context"
	"main/internal/model"
)

type AuthRepository interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, info *model.UserInfo) (string, error)
}
