package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unkabas/JWTGo/config"
	"github.com/unkabas/JWTGo/controllers"
)

func main() {
	config.LoadEnvs()
	config.ConnectDB()
	r := gin.Default()
	r.POST("/reg", controllers.Registration)
	r.POST("/login", controllers.Login)
	r.Run(":8080")
}
