package unit_test

import (
	"os"
	"testing"

	"stadium-builder-backend/config"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestConnectDatabase(t *testing.T) {
	// Load environment variables from .env file
	err := godotenv.Load("../.env")
	assert.NoError(t, err, "Failed to load .env file")

	// Ensure TEST_DATABASE_URL is set
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		t.Fatal("TEST_DATABASE_URL is not set")
	}

	// Connect to the database
	db, err := config.ConnectDatabase()

	// Assertions
	assert.NoError(t, err, "Database connection should succeed")
	assert.NotNil(t, db, "DB instance should not be nil")
}
