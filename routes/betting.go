package routes

import (
	"log"
	"net/http"
	"stadium-builder-backend/services"

	"github.com/gin-gonic/gin"
)

func BettingRoutes(router *gin.Engine) {
	router.GET("/betting", func(c *gin.Context) {
		games, err := services.FetchBettingData()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		log.Println("Fetched Games:", games)
		c.JSON(http.StatusOK, gin.H{"Games": games})
	
		if err := services.SaveBettingData(games); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{"message": "Data fetched and saved successfully"})
	})
	
}
