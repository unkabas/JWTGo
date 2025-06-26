package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unkabas/JWTGo/config"
	"github.com/unkabas/JWTGo/controllers"
	"github.com/unkabas/JWTGo/middleware"
)

func main() {
	config.LoadEnvs()
	config.ConnectDB()
	r := gin.Default()
	r.POST("/reg", controllers.Registration)
	r.POST("/login", controllers.Login)
	r.GET("/hello", middleware.AuthMiddleware, controllers.SayHello)
	r.Run(":8080")
}
