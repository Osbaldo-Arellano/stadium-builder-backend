package unit_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"stadium-builder-backend/services"

	"github.com/stretchr/testify/assert"
)

func TestFetchBettingData(t *testing.T) {
	// Mock server to simulate API response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulated API response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[
			{
				"id": "1",
				"sport_key": "soccer",
				"sport_title": "Soccer",
				"home_team": "Team A",
				"away_team": "Team B",
				"commence_time": "2024-12-20T18:00:00Z",
				"bookmakers": []
			}
		]`))
	}))
	defer mockServer.Close()

	// Set the BETTING_API_KEY for the test
	os.Setenv("BETTING_API_KEY", "test_api_key")

	// Call FetchBettingData with the mock server's URL
	games, err := services.FetchBettingData(mockServer.URL)

	// Assertions
	assert.NoError(t, err, "Fetching betting data should not return an error")
	assert.NotNil(t, games, "Games data should not be nil")
	assert.Equal(t, 1, len(games), "Games list should contain one entry")
	assert.Equal(t, "1", games[0].ID, "Game ID should match the mock response")
	assert.Equal(t, "soccer", games[0].SportKey, "SportKey should match the mock response")
	assert.Equal(t, "Team A", games[0].HomeTeam, "HomeTeam should match the mock response")
	assert.Equal(t, "Team B", games[0].AwayTeam, "AwayTeam should match the mock response")
}
