package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/unkabas/JWTGo/config"
	"github.com/unkabas/JWTGo/models"
	"os"
	"strings"
)

func AddExpense(c *gin.Context) {
	var newExpense models.Expense
	if err := c.ShouldBindJSON(&newExpense); err != nil {
		c.JSON(400, gin.H{
			"message": "Not a JSON",
		})
		return
	}
	if newExpense.Name == "" || newExpense.Price == 0 || newExpense.Date == "" {
		c.JSON(400, gin.H{
			"message": "No Name or Price or Date added",
		})
		return
	}

	secret := os.Getenv("SECRET")
	tokenString := c.GetHeader("Authorization")
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	} else {
		c.JSON(401, gin.H{"message": "Invalid Authorization header format"})
		return
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
		res := claims["sub"].(string)
		newExpense.Author = res
		fmt.Println(newExpense.Author)
	} else {
		fmt.Println("Invalid token")
	}

	if err := config.DB.Create(&newExpense).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "Expense not created",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success",
	})
}
