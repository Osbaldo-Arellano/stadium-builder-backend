package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Fatal("Missing REDIS_URL environment variable")
	}

	// Init Redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisURL, // e.g., "localhost:6379"
		Password: "", // Optional password
		DB:       0,                           // Use default DB
	})

	// Test the connection with a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := RedisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis at %s: %v", redisURL, err)
	}

	log.Println("Connected to Redis successfully")
}

// CloseRedis gracefully closes the Redis client
func CloseRedis() {
	if err := RedisClient.Close(); err != nil {
		log.Printf("Error closing Redis connection: %v", err)
	} else {
		log.Println("Redis connection closed successfully")
	}
}
