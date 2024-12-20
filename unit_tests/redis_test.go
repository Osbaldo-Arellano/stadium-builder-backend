package unit_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"stadium-builder-backend/config"
	"stadium-builder-backend/models"
	"stadium-builder-backend/services"

	"github.com/go-co-op/gocron"
	"github.com/stretchr/testify/assert"
)

func TestCacheBettingData(t *testing.T) {
	// Initialize a mock Redis connection
	config.ConnectRedis()
	defer config.RedisClient.FlushAll(context.Background()) // Clear Redis after test

	// Mock data
	mockGames := []models.APIResponseGame{
		{ID: "1", SportKey: "soccer", HomeTeam: "Team A", AwayTeam: "Team B"},
	}

	// Cache the data
	err := services.CacheBettingData(mockGames)
	assert.NoError(t, err, "Caching betting data should not return an error")

	// Verify the data is cached
	cachedData, err := config.RedisClient.Get(context.Background(), "betting_data").Result()
	assert.NoError(t, err, "Retrieving cached data should not return an error")

	var cachedGames []models.APIResponseGame
	err = json.Unmarshal([]byte(cachedData), &cachedGames)
	assert.NoError(t, err, "Unmarshaling cached data should not return an error")
	assert.Equal(t, mockGames, cachedGames, "Cached data should match the input data")
}

func TestSchedulerAndCacheIntegration(t *testing.T) {
	// Initialize Redis
	config.ConnectRedis()
	defer config.RedisClient.FlushAll(context.Background()) // Clean up Redis after test

	// Step 1: Set initial mock data in Redis
	initialMockData := []models.APIResponseGame{
		{
			ID:       "1",
			SportKey: "soccer",
			HomeTeam: "Initial Team A",
			AwayTeam: "Initial Team B",
		},
	}
	initialMockJSON, _ := json.Marshal(initialMockData)
	err := config.RedisClient.Set(context.Background(), "betting_data", initialMockJSON, 5*time.Minute).Err()
	assert.NoError(t, err, "Setting initial data in Redis should not return an error")

	// Step 2: Create a mock server for the scheduler to fetch new data
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[
			{
				"id": "2",
				"sport_key": "soccer",
				"home_team": "Updated Team A",
				"away_team": "Updated Team B"
			}
		]`))
	}))
	defer mockServer.Close()

	// Step 3: Set up the scheduler to fetch data from the mock server
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(1).Seconds().Do(func() {
		_, err := services.FetchBettingData(mockServer.URL, true)
		assert.NoError(t, err, "Scheduler should fetch and update cache without errors")
	})
	scheduler.StartAsync()
	defer scheduler.Stop()

	// Step 4: Wait for the scheduler to update the cache
	time.Sleep(5 * time.Second) // Wait for at least one scheduler cycle

	// Step 5: Verify the cache has been updated with new data
	cachedData, err := services.GetCachedBettingData()
	assert.NoError(t, err, "Fetching cached data should not return an error")
	assert.Equal(t, 1, len(cachedData), "Cache should contain one updated entry")
	assert.Equal(t, "Updated Team A", cachedData[0].HomeTeam, "Home team should match updated mock data")
	assert.Equal(t, "Updated Team B", cachedData[0].AwayTeam, "Away team should match updated mock data")
}

