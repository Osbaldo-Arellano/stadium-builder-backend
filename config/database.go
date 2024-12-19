package config

import (
	"log"
	"os"
	"stadium-builder-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Fetch the database connection string from environment variables
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Missing DATABASE_URL environment variable")
	}

	// Connect to the PostgreSQL database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := database.AutoMigrate(&models.Leaderboard{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	if err := database.AutoMigrate(&models.Game{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	DB = database
	log.Println("Database connected and migrations applied successfully")
}
