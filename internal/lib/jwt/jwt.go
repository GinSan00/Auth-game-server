package jwt

import (
	"main/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(user model.User, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["id"] = user.ID
	claims["email"] = user.Info.Email
	claims["nickname"] = user.Info.Nickname
	claims["elo"] = user.Info.Elo

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
