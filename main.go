package main

import (
	"log"
	"os"

	"stadium-builder-backend/config"
	"stadium-builder-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	// Connect to Postgres DB
	config.ConnectDatabase()

	// Initialize the Gin router
	router := gin.Default()
	
	// Register routes
	routes.HealthRoutes(router)
	routes.BettingRoutes(router)
	routes.LeaderboardRoutes(router)



	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
