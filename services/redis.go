package services

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"stadium-builder-backend/config"
	"stadium-builder-backend/models"

	"github.com/redis/go-redis/v9"
)

// Cache key for betting data
const BettingDataKey = "betting_data"

// GetCachedBettingData fetches betting data from Redis
func GetCachedBettingData() ([]models.APIResponseGame, error) {
	ctx := context.Background()

	// Attempt to get data from Redis
	val, err := config.RedisClient.Get(ctx, BettingDataKey).Result()
	if err != nil {
		// Handle key not found separately
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	// Unmarshal JSON data into Go struct
	var games []models.APIResponseGame
	if err := json.Unmarshal([]byte(val), &games); err != nil {
		return nil, err
	}

	return games, nil
}

// CacheBettingData saves betting data to Redis
func CacheBettingData(data []models.APIResponseGame) error {
	ctx := context.Background()

	// Marshal data into JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Cache data with expiration
	expiration := 5 * time.Minute
	return config.RedisClient.Set(ctx, BettingDataKey, jsonData, expiration).Err()
}
