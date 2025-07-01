package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/unkabas/JWTGo/config"
	"github.com/unkabas/JWTGo/models"
	"github.com/unkabas/JWTGo/services"
)

func Refresh(c *gin.Context) {

	var inputToken models.InputAuth
	var session models.Session

	if err := c.ShouldBindJSON(&inputToken); err != nil {
		c.JSON(400, gin.H{"message": "Not a JSON"})
		return
	}
	if inputToken.Refresh == "" {
		c.JSON(400, gin.H{"message": "No refresh token"})
		return
	}

	username := services.DecodeJwt(c)

	if err := config.DB.Where("username = ? AND refresh_token = ?", username, inputToken.Refresh).First(&session).Error; err != nil {
		c.JSON(400, gin.H{"message": "No such session"})
		return
	}

	tokenString, err := services.SetJWT(username)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cant generate Token"})
		return
	}

	c.JSON(200, gin.H{
		"message":  "success",
		"username": username,
		"new AT":   tokenString,
	})

}
