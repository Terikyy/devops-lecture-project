package jwt

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func VerifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})
	return err == nil && token.Valid
}

func ExtractBearerToken(authHeader string) string {
	return strings.TrimPrefix(authHeader, "Bearer ")
}
