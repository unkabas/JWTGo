package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unkabas/JWTGo/services"
	"log"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	res := c.GetHeader("Authorization")
	tokenString := strings.Split(res, " ")
	if tokenString[0] != "Bearer" {
		c.JSON(400, gin.H{
			"error": "token is incorrect type",
		})
		c.Abort()
	}
	token := tokenString[1]
	if token == "" {
		c.JSON(400, gin.H{
			"error": "token is null",
		})
		c.Abort()
	}
	if token == "" {
		c.JSON(400, gin.H{
			"error": "token is null",
		})
		c.Abort()

	}
	if err := services.VerifyToken(token); err != nil {
		c.JSON(400, gin.H{
			"error": "token is not valid",
		})
		log.Println(err)
		c.Abort()
		return
	}
	c.Next()
}
