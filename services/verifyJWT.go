package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func VerifyToken(tokenString string) error {
	secret := os.Getenv("SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
