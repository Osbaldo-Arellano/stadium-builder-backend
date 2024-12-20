package unit_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"stadium-builder-backend/config"
	"stadium-builder-backend/services"

	"github.com/stretchr/testify/assert"
)

func TestFetchBettingData(t *testing.T) {
	// Set environment variables
	os.Setenv("REDIS_URL", "localhost:6379") // Ensure Redis URL is set before initialization
	os.Setenv("BETTING_API_KEY", "test_api_key")

	// Initialize Redis
	config.ConnectRedis()
	defer config.RedisClient.FlushAll(context.Background()) // Clean up Redis after test

	// Mock server to simulate external API response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[
			{
				"id": "1",
				"sport_key": "soccer",
				"home_team": "Team A",
				"away_team": "Team B"
			}
		]`))
	}))
	defer mockServer.Close()

	// Call FetchBettingData
	games, err := services.FetchBettingData(mockServer.URL, true)
	assert.NoError(t, err, "Fetching betting data should not return an error")
	assert.Equal(t, 1, len(games), "Should return one game")
	assert.Equal(t, "Team A", games[0].HomeTeam, "Home team should match mock data")
}
