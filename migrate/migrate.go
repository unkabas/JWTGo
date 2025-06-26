package main

import (
	"github.com/unkabas/JWTGo/config"
	"github.com/unkabas/JWTGo/models"
)

func main() {
	config.LoadEnvs()
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Expense{})
}
