package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unkabas/JWTGo/services"
	"log"
)

func AuthMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(400, gin.H{
			"error": "token is not defined",
		})
		log.Println(err)
		c.Abort()
		return
	}
	if tokenString == "" {
		c.JSON(400, gin.H{
			"error": "token is null",
		})
		c.Abort()

	}
	if err := services.VerifyToken(tokenString); err != nil {
		c.JSON(400, gin.H{
			"error": "token is not valid",
		})
		log.Println(err)
		c.Abort()
		return
	}
	c.Next()
}
