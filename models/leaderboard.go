package models

import "gorm.io/gorm"

type Leaderboard struct {
	gorm.Model
	PlayerID   string  `gorm:"uniqueIndex"` 
	PlayerName string  
	Score      int     
	Rank       int     
}
