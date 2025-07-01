package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unkabas/JWTGo/config"
	"github.com/unkabas/JWTGo/models"
	"github.com/unkabas/JWTGo/services"
)

func Login(c *gin.Context) {
	var input models.InputAuth

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"message": "Not a JSON"})
		return
	}
	if input.Username == "" || input.Password == "" {
		c.JSON(400, gin.H{"message": "No username or password added"})
		return
	}

	var userFound models.User
	if err := config.DB.Where("username=?", input.Username).First(&userFound).Error; err != nil {
		c.JSON(401, gin.H{"message": "Incorrect username, try to sign in"})
		return
	}

	match := services.CheckPasswordHash(input.Password, userFound.Password)
	if !match {
		c.JSON(401, gin.H{"message": "Incorrect username or password"})
		return
	}

	tokenString, err := services.SetJWT(userFound.Username)
	if err != nil {
		c.JSON(400, gin.H{"error": "Cant generate Token"})
		return
	}

	refresh, err := services.SetRefresh()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	session := models.Session{
		Username:     input.Username,
		RefreshToken: refresh,
	}
	if err := config.DB.Create(&session).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create refresh",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success login",
		"token":   tokenString,
		"refresh": refresh,
	})
}
