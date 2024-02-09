package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtCustomClaims struct {
	UserId int    `json:"userId"`
	Name   string `json:"bool"`
	jwt.RegisteredClaims
}

func GenerateTokenJWT(userid int, name string) string {
	var claims = jwtCustomClaims{
		userid,
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	resultJWT, _ := token.SignedString([]byte("123"))

	return resultJWT
}
