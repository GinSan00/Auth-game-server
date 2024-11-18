package auth

import (
	"main/internal/service"

	desc "github.com/GinSan00/auth-game-server/pkg/auth/authv1"
)

type Implementation struct {
	authService service.AuthService
	desc.UnimplementedAuthServer
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{authService: authService}
}
