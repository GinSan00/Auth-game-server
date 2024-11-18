package auth

import (
	"context"

	desc "github.com/GinSan00/auth-game-server/pkg/auth/authv1"
)

func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	token, err := i.authService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &desc.LoginResponse{
		Token: token,
	}, nil
}
