package main

import (
	"log"
	"os"
	"time"

	"stadium-builder-backend/config"
	"stadium-builder-backend/routes"
	"stadium-builder-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func getEnvVar(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		if defaultValue == "" {
			log.Fatalf("Environment variable %s is required but not set", key)
		}
		return defaultValue
	}
	return value
}

func main() {
	// Connect to Postgres DB
	if _, err := config.ConnectDatabase(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize the Gin router
	router := gin.Default()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Initialize Redis connection
	config.ConnectRedis()

	// Betting odds URL
	apiURL := getEnvVar("ODDS_API_URL", "")

	// Set up scheduler
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minutes().Do(func() {
		games, err := services.FetchBettingData(apiURL, false)
		if err != nil {
			log.Printf("Failed to fetch betting data: %v", err)
		} else {
			if err := services.CacheBettingData(games); err != nil {
				log.Printf("Failed to cache betting data: %v", err)
			} else {
				log.Println("Betting data updated and cached successfully")
			}
		}
	})
	s.StartAsync()
	defer cleanup(s)

	// Register routes
	routes.HealthRoutes(router)
	routes.BettingRoutes(router)
	routes.LeaderboardRoutes(router)

	// Start the server
	port := getEnvVar("PORT", "8080")
	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func cleanup(scheduler *gocron.Scheduler) {
	scheduler.Stop()
	config.RedisClient.Close()
	log.Println("Cleaned up resources")
}
