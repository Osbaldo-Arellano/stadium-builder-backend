package unit_test

import (
	"testing"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/stretchr/testify/assert"
)

func TestSchedulerPeriodicUpdate(t *testing.T) {
	// Mock a scheduler
	s := gocron.NewScheduler(time.UTC)
	mockExecuted := false

	// Schedule a mock task
	s.Every(1).Seconds().Do(func() {
		mockExecuted = true
	})

	// Start the scheduler
	s.StartAsync()
	defer s.Stop()

	// Wait for the task to execute
	time.Sleep(2 * time.Second)

	assert.True(t, mockExecuted, "Scheduled task should execute at least once")
}
