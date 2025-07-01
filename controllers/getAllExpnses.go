package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/unkabas/JWTGo/config"
	"github.com/unkabas/JWTGo/models"
	"github.com/unkabas/JWTGo/services"
)

func GetAllExpnses(c *gin.Context) {
	user := services.DecodeJwt(c)
	var expnses []models.Expense
	sortBy := c.DefaultQuery("sort", "")
	query := config.DB.Where("author = ?", user)

	switch sortBy {
	case "price_asc":
		query = config.DB.Order("price ASC")
	case "price_desc":
		query = config.DB.Order("price DESC")
	case "date_asc":
		query = config.DB.Order("date ASC")
	case "date_desc":
		query = config.DB.Order("date DESC")
	}
	if err := query.Find(&expnses).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "no expenses found",
		})
		return
	}
	c.JSON(200, gin.H{
		"expenses": expnses,
	})
}
