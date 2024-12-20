package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"stadium-builder-backend/models"
)

// FetchBettingData retrieves betting data from the API
func FetchBettingData(apiURL string) ([]models.APIResponseGame, error) {
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
