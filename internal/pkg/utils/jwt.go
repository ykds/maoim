package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var (
	signKey = []byte("maoim")
)

type Claims struct {
	jwt.StandardClaims
	ID       string
	Username string
}

func GenToken(ID string, username string) (string, error) {
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
		},
		ID:       ID,
		Username: username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signKey)
}

func ValidToken(token string) (string, string, error) {
	claims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		return "", "", err
	}
	if c, ok := claims.Claims.(*Claims); ok && claims.Valid {
		return c.ID, c.Username, nil
	}
	return "", "", fmt.Errorf("token invalid")
}
