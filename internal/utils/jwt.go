package utils

import (
	"time"

	"github.com/amrit713/food-delivery/config"
	"github.com/golang-jwt/jwt/v5"
)

var JwtSecrect = []byte(config.LoadEnv("JWT_SECRET", "c1vpTdSvxX"))

func GenerateJWT(userID string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JwtSecrect)
}
