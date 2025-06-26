package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/unkabas/JWTGo/config"
	"github.com/unkabas/JWTGo/models"
	"github.com/unkabas/JWTGo/services"
)

func Registration(c *gin.Context) {
	var input models.InputAuth
	var existingUser models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"message": "Not a JSON",
		})
		return
	}
	if input.Username == "" || input.Password == "" {
		c.JSON(400, gin.H{
			"message": "Username or Email or Password are not filled",
		})
		return
	}
	hash, _ := services.HashPassword(input.Password)
	if err := config.DB.Where("username = ?", input.Username).First(&existingUser).Error; err != nil {
		user := models.User{
			Username: input.Username,
			Password: hash,
		}
		if err := config.DB.Create(&user).Error; err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to create user",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Success",
		})

	} else {
		c.JSON(409, gin.H{
			"message": "Username already taken",
		})
	}
}
