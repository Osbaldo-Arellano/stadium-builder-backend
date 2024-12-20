package unit_test

import (
	"os"
	"testing"

	"stadium-builder-backend/config"
	"stadium-builder-backend/models"
	"stadium-builder-backend/services"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestUpdatePlayerScore(t *testing.T) {
	// Load .env file only in local environments
	if os.Getenv("CI") == "" {
		err := godotenv.Load("../.env")
		assert.NoError(t, err, "Failed to load .env file")
	}

	// Ensure TEST_DATABASE_URL is set
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		t.Fatal("TEST_DATABASE_URL is not set")
	}

	// Connect to the database
	db, err := config.ConnectDatabase()
	assert.NoError(t, err, "Database connection should succeed")
	assert.NotNil(t, db, "Database instance should not be nil")

	config.DB.Exec("DELETE FROM leaderboards") // Clear existing data

	// Add a new player score
	err = services.UpdatePlayerScore("player1", "John", 100)
	assert.NoError(t, err, "Adding player score should not return an error")

	// Verify the player score is added
	var leaderboard models.Leaderboard
	config.DB.Where("player_id = ?", "player1").First(&leaderboard)
	assert.Equal(t, "John", leaderboard.PlayerName, "Player name should match")
	assert.Equal(t, 100, leaderboard.Score, "Score should match")

	// Update the player's score
	err = services.UpdatePlayerScore("player1", "John", 50)
	assert.NoError(t, err, "Updating player score should not return an error")

	// Verify the updated score
	config.DB.Where("player_id = ?", "player1").First(&leaderboard)
	assert.Equal(t, 150, leaderboard.Score, "Score should be updated correctly")
}

func TestGetTopPlayers(t *testing.T) {
	// Load .env file only in local environments
	if os.Getenv("CI") == "" {
		err := godotenv.Load("../.env")
		assert.NoError(t, err, "Failed to load .env file")
	}

	// Ensure TEST_DATABASE_URL is set
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		t.Fatal("TEST_DATABASE_URL is not set")
	}

	// Connect to the database
	db, err := config.ConnectDatabase()
	assert.NoError(t, err, "Database connection should succeed")
	assert.NotNil(t, db, "Database instance should not be nil")

	config.DB.Exec("DELETE FROM leaderboards") // Clear existing data

	// Add multiple players
	services.UpdatePlayerScore("player1", "John", 100)
	services.UpdatePlayerScore("player2", "Alice", 200)
	services.UpdatePlayerScore("player3", "Bob", 50)

	// Fetch top players
	topPlayers, err := services.GetTopPlayers(2)
	assert.NoError(t, err, "Fetching top players should not return an error")
	assert.Equal(t, 2, len(topPlayers), "Should return the correct number of players")
	assert.Equal(t, "Alice", topPlayers[0].PlayerName, "Top player should be Alice")
	assert.Equal(t, "John", topPlayers[1].PlayerName, "Second player should be John")
}
