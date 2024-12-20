package config

import (
	"fmt"
	"log"
	"os"
	"stadium-builder-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	// Use TEST_DATABASE_URL if running tests, otherwise DATABASE_URL
	dsn := os.Getenv("TEST_DATABASE_URL")
	if os.Getenv("GO_ENV") != "test" { // Use DATABASE_URL if not in test mode
		dsn = os.Getenv("DATABASE_URL")
	}

	if dsn == "" {
		return nil, fmt.Errorf("missing database connection string")
	}

	// Set logger mode: Silent during tests, Default otherwise
	logMode := logger.Default
	if os.Getenv("GO_ENV") == "test" {
		logMode = logger.Default.LogMode(logger.Silent)
	}

	// Connect to the database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logMode,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Run migrations
	if err := database.AutoMigrate(&models.Leaderboard{}, &models.Game{}); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %v", err)
	}

	DB = database
	log.Println("Database connected successfully")

	return database, nil
}
