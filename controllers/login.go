package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/unkabas/JWTGo/config"
	"github.com/unkabas/JWTGo/models"
	"github.com/unkabas/JWTGo/services"
)

func Login(c *gin.Context) {
	var input models.InputAuth
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"message": "Not a JSON",
		})
		return
	}
	if input.Username == "" || input.Password == "" {
		c.JSON(400, gin.H{
			"message": "No username or password added",
		})
		return
	}
	var userFound models.User
	if err := config.DB.Where("username=? ", input.Username).First(&userFound).Error; err != nil {
		c.JSON(401, gin.H{
			"message": "Incorrect username, try to sign in",
		})
	} else {
		match := services.CheckPasswordHash(input.Password, userFound.Password)
		if !match {
			c.JSON(401, gin.H{
				"message": "Incorrect username or password",
			})
		} else {
			c.JSON(200, gin.H{
				"message":  "success",
				"username": input.Username,
			})
		}
	}
}
