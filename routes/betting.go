package routes

import (
	"net/http"
	"os"
	"stadium-builder-backend/services"

	"github.com/gin-gonic/gin"
)

func BettingRoutes(router *gin.Engine) {
	router.GET("/betting", func(c *gin.Context) {
		// Validate API URL
		apiURL := os.Getenv("ODDS_API_URL")
		if apiURL == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Betting API URL not configured"})
			return
		}

		// Attempt to retrieve data from Redis cache
		cachedData, err := services.GetCachedBettingData()
		if err == nil && cachedData != nil {
			c.JSON(http.StatusOK, gin.H{
				"Games": cachedData,
				"Source": "cache",
			})
			return
		}

		// Fetch betting data from the external API
		games, err := services.FetchBettingData(apiURL, false)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch betting data from API"})
			return
		}

		// Cache the fetched data in Redis
		err = services.CacheBettingData(games)
		if err != nil {
			// Log the error but still return the fetched data to the client
			c.JSON(http.StatusOK, gin.H{
				"Games":  games,
				"Source": "api",
				"Warning": "Data fetched but caching failed",
			})
			return
		}

		// Return the fetched data
		c.JSON(http.StatusOK, gin.H{
			"Games":  games,
			"Source": "api",
		})
	})
}
