package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/unkabas/JWTGo/config"
	"github.com/unkabas/JWTGo/models"
	"github.com/unkabas/JWTGo/services"
)

func DeleteExpense(c *gin.Context) {
	user := services.DecodeJwt(c)
	name := c.Param("name")
	var deleteExpense models.Expense
	config.DB.Where("author = ? AND name = ?", user, name).First(&deleteExpense)
	if err := config.DB.Delete(&deleteExpense).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "Expense not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Success",
	})
}
