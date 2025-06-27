package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/unkabas/JWTGo/config"
	"github.com/unkabas/JWTGo/models"
	"github.com/unkabas/JWTGo/services"
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
	res := services.DecodeJwt(c)
	newExpense.Author = res
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
