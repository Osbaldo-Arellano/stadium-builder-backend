package routes

import (
	"github.com/gin-gonic/gin"
)

func HealthRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "API is running",
		})
	})
}
