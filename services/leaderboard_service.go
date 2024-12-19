package services

import (
	"fmt"
	"stadium-builder-backend/config"
	"stadium-builder-backend/models"
)

// Add a player's score to the leaderboard or update their existing score
func UpdatePlayerScore(playerID, playerName string, score int) error {
	var leaderboard models.Leaderboard

	// Check if player already exists
	if err := config.DB.Where("player_id = ?", playerID).First(&leaderboard).Error; err != nil {
		if err.Error() == "record not found" {
			// If not found, create a new record
			newEntry := models.Leaderboard{
				PlayerID:   playerID,
				PlayerName: playerName,
				Score:      score,
			}
			if err := config.DB.Create(&newEntry).Error; err != nil {
				return fmt.Errorf("failed to create leaderboard entry: %v", err)
			}
		} else {
			return fmt.Errorf("failed to query leaderboard: %v", err)
		}
	} else {
		// If found, update the score
		leaderboard.Score += score
		if err := config.DB.Save(&leaderboard).Error; err != nil {
			return fmt.Errorf("failed to update leaderboard entry: %v", err)
		}
	}

	return nil
}

// Get the top players in the leaderboard
func GetTopPlayers(limit int) ([]models.Leaderboard, error) {
	var leaderboard []models.Leaderboard

	// Fetch players ordered by score (descending)
	if err := config.DB.Order("score desc").Limit(limit).Find(&leaderboard).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch leaderboard: %v", err)
	}

	// Assign ranks based on the order
	for i := range leaderboard {
		leaderboard[i].Rank = i + 1
	}

	return leaderboard, nil
}

