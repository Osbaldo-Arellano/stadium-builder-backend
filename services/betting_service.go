package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"stadium-builder-backend/config"
	"stadium-builder-backend/models"
)

type Game struct {
	ID        string  `json:"id"`
	HomeTeam  string  `json:"home_team"`
	AwayTeam  string  `json:"away_team"`
	StartTime string  `json:"start_time"`
	Odds      float64 `json:"odds"`
}

func FetchBettingData() ([]Game, error) {
	apiKey := os.Getenv("BETTING_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("missing API key")
	}

	url := "https://api.the-odds-api.com/v4/sports/soccer/odds?apiKey=" + apiKey + "&regions=us&markets=h2h"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch betting data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read API response: %v", err)
	}

	var games []Game
	if err := json.Unmarshal(body, &games); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %v", err)
	}

	return games, nil
}

func SaveBettingData(games []Game) error {
	for _, game := range games {
		dbGame := models.Game{
			ExternalID: game.ID,
			HomeTeam:   game.HomeTeam,
			AwayTeam:   game.AwayTeam,
			StartTime:  game.StartTime,
			Odds:       game.Odds,
		}

		// Use GORM's Create or Update logic
		if err := config.DB.Where("external_id = ?", game.ID).FirstOrCreate(&dbGame).Error; err != nil {
			return fmt.Errorf("failed to save game: %v", err)
		}
	}
	return nil
}
