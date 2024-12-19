package routes

import (
	"net/http"
	"stadium-builder-backend/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LeaderboardRoutes(router *gin.Engine) {
	// Add or update a player's score
	router.POST("/leaderboard", func(c *gin.Context) {
		playerID := c.PostForm("player_id")
		playerName := c.PostForm("player_name")
		score, err := strconv.Atoi(c.PostForm("score"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid score"})
			return
		}

		if err := services.UpdatePlayerScore(playerID, playerName, score); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Player score updated successfully"})
	})

	// Get the top players
	router.GET("/leaderboard", func(c *gin.Context) {
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}

		players, err := services.GetTopPlayers(limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, players)
	})
}
