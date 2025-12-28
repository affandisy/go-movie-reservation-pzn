package utils

import (
	"go-movie-reservation/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWTToken(userID string, role string) (string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	parseJWT, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(cfg.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return parseJWT, nil
}
