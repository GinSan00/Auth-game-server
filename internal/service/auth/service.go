package auth

import (
	"main/internal/repository"
	def "main/internal/service"
)

var _ def.AuthService = (*service)(nil)

type service struct {
	authRepository repository.AuthRepository
}

func NewService(authRepository repository.AuthRepository) *service {
	return &service{authRepository: authRepository}
}
