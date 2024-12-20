package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"stadium-builder-backend/config"
	"stadium-builder-backend/models"
	"time"
)

const BettingCacheKey = "betting_data"

// FetchBettingData retrieves betting data from the API or cache
func FetchBettingData(apiURL string, forceRefresh bool) ([]models.APIResponseGame, error) {
    if !forceRefresh {
        cachedGames, err := fetchFromCache(BettingCacheKey)
        if err == nil {
            return cachedGames, nil
        }
    }
    // If forceRefresh is true or cache miss occurred:
    apiGames, err := fetchFromAPI(apiURL)
    if err != nil {
        return nil, err
    }
    if err := cacheData(BettingCacheKey, apiGames); err != nil {
        return nil, fmt.Errorf("failed to cache API response: %v", err)
    }
    return apiGames, nil
}


// fetchFromCache retrieves betting data from Redis
func fetchFromCache(cacheKey string) ([]models.APIResponseGame, error) {
	ctx := context.Background()

	// Get data from Redis
	cachedData, err := config.RedisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		return nil, err // Cache miss or Redis error
	}

	// Unmarshal JSON data
	var games []models.APIResponseGame
	if err := json.Unmarshal([]byte(cachedData), &games); err != nil {
		return nil, fmt.Errorf("failed to parse cached data: %v", err)
	}

	return games, nil
}

// fetchFromAPI retrieves betting data directly from the API
func fetchFromAPI(apiURL string) ([]models.APIResponseGame, error) {
	apiKey := os.Getenv("BETTING_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("missing API key")
	}

	// Parse the API URL and add the apiKey as a query parameter
	parsedURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, fmt.Errorf("invalid API URL: %v", err)
	}
	query := parsedURL.Query()
	query.Set("apiKey", apiKey)
	parsedURL.RawQuery = query.Encode()

	// Make the HTTP request
	resp, err := http.Get(parsedURL.String())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch betting data: %v", err)
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	// Read and parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read API response: %v", err)
	}

	var apiGames []models.APIResponseGame
	if err := json.Unmarshal(body, &apiGames); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %v", err)
	}

	return apiGames, nil
}

// cacheData stores betting data in Redis
func cacheData(cacheKey string, data []models.APIResponseGame) error {
	ctx := context.Background()

	// Marshal data into JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data for caching: %v", err)
	}

	// Cache data with expiration (5 minutes)
	return config.RedisClient.Set(ctx, cacheKey, jsonData, 5*time.Minute).Err()
}
