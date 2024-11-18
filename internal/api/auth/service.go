package auth

import (
	"main/internal/service"

	desc "github.com/ginsan00/auth-game-server/pkg/auth"
)

type Implementation struct {
	authService service.AuthService
	desc.UnimplementedAuthServer
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{authService: authService}

}
