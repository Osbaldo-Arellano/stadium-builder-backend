package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	ExternalID string `gorm:"uniqueIndex"`
	HomeTeam   string
	AwayTeam   string
	StartTime  string
	Odds       float64
}
