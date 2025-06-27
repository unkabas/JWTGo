package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

func DecodeJwt(c *gin.Context) string {
	secret := os.Getenv("SECRET")
	tokenString := c.GetHeader("Authorization")
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	} else {
		c.JSON(401, gin.H{
			"message": "Invalid Authorization header format",
		})
		return ""
	}
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		c.JSON(400, gin.H{
			"message": "no token",
		})
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub, ok := claims["sub"].(string)
		if !ok {
			c.JSON(400, gin.H{"message": "Invalid sub claim"})
			return ""
		}
		fmt.Println("JWT sub claim:", sub)
		return sub
	}
	c.JSON(401, gin.H{"message": "Invalid token"})
	return ""
}
