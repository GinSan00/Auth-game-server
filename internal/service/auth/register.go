package auth

import (
	"context"
	"log"
	"main/internal/model"

	"github.com/google/uuid"
)

func (s *service) Register(ctx context.Context, info *model.UserInfo) (string, error) {
	userUUID, err := uuid.NewUUID()
	if err != nil {
		log.Printf("ошибка получения пользователя: %v\n", err)
		return "", err
	}

	err = s.authRepository.Register(ctx, userUUID.String(), info)
	if err != nil {
		log.Printf("ошибка получения пользователя: %v\n", err)
		return "", err
	}

	return userUUID.String(), nil

}
