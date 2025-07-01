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
	{
		auth := r.Group("/auth")
		auth.POST("/reg", controllers.Registration)
		auth.POST("/login", controllers.Login)
		auth.POST("/refresh", middleware.AuthMiddleware, controllers.Refresh)
	}
	{
		expense := r.Group("/expense")
		expense.POST("/add", middleware.AuthMiddleware, controllers.AddExpense)
		expense.GET(":name/delete", middleware.AuthMiddleware, controllers.DeleteExpense)
		expense.GET("/all", middleware.AuthMiddleware, controllers.GetAllExpnses)
	}

	r.GET("/sayHello", middleware.AuthMiddleware, controllers.SayHello)

	r.Run(":8080")
}
