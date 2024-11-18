package auth

import (
	"context"
	"log"
	"main/internal/model"
)

func (s *service) Login(ctx context.Context, email, password string) (string, error) {
	token, err := s.authRepository.Login(ctx, email, password)
	if err != nil {
		return "", err
	}
	if token == "" {
		log.Printf("User not found: %s", email)
		return "", model.ErrorUserNotFound
	}

	return token, nil
}
