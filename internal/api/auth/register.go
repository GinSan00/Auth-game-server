package auth

import (
	"context"

	desc "github.com/GinSan00/auth-game-server/pkg/auth/authv1"
)

func (i *Implementation) Register(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	uID, err := i.authService.Register(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &desc.RegisterResponse{
		user_id: uID,
	}, nil
}
