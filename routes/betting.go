package routes

import (
	"net/http"
	"stadium-builder-backend/services"

	"github.com/gin-gonic/gin"
)

func BettingRoutes(router *gin.Engine) {
	router.GET("/betting", func(c *gin.Context) {
		// Fetch betting data
		apiURL := "https://api.the-odds-api.com/v4/sports/americanfootball_nfl/odds?regions=us&oddsFormat=american&apiKey="
		games, err := services.FetchBettingData(apiURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Send the games data to the client
		c.JSON(http.StatusOK, gin.H{"Games": games})
	})
}
