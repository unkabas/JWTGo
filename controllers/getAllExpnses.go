package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/unkabas/JWTGo/config"
	"github.com/unkabas/JWTGo/models"
	"github.com/unkabas/JWTGo/services"
)

func GetAllExpnses(c *gin.Context) {
	user := services.DecodeJwt(c)
	var allExpnses []models.Expense

	if err := config.DB.Where("author = ?", user).Find(&allExpnses).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "no expnses found",
		})
		return
	}
	c.JSON(200, gin.H{
		"expnses": allExpnses,
	})
}
