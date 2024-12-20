package models

import (
	"gorm.io/gorm"
)

type Outcomes struct {
	gorm.Model
	Name      string
	Price     int64
	MarketID  uint `gorm:"index"` // Foreign key to Markets
}

type Markets struct {
	gorm.Model
	Key        string
	LastUpdate string
	Outcome    []Outcomes `gorm:"foreignKey:MarketID"` // Define relationship to Outcomes
	BookmakerID uint `gorm:"index"` // Foreign key to Bookmaker
}

type Bookmaker struct {
	gorm.Model
	Key        string
	LastUpdate string
	Markets    []Markets `gorm:"foreignKey:BookmakerID"` // Define relationship to Markets
	GameID     uint `gorm:"index"` // Foreign key to Game
}

type Game struct {
	gorm.Model
	ExternalID   string `gorm:"uniqueIndex"`
	SportsKey    string
	CommenceTime string
	HomeTeam     string
	AwayTeam     string
	Bookmakers   []Bookmaker `gorm:"foreignKey:GameID"` // Define relationship to Bookmakers
}


// Example:
// {
//   "ExternalID": "a9bf9c4e9be5107617933461fa1f7f382",
//   "SportsKey": "americanfootball_nfl",
//   "CommenceTime": "2024-12-20T01:16:00Z",
//   "HomeTeam": "Los Angeles Chargers",
//   "AwayTeam": "Denver Broncos",
//   "Bookmakers": [
//     {
//       "Key": "fanduel",
//       "LastUpdate": "2024-12-19T18:41:04Z",
//       "Markets": [
//         {
//           "Key": "h2h",
//           "LastUpdate": "2024-12-19T18:41:04Z",
//           "Outcome": [
//             {"Name": "Denver Broncos", "Price": 126},
//             {"Name": "Los Angeles Chargers", "Price": -148}
//           ]
//         }
//       ]
//     },
//     {
//       "Key": "draftkings",
//       "LastUpdate": "2024-12-19T18:40:45Z",
//       "Markets": [
//         {
//           "Key": "h2h",
//           "LastUpdate": "2024-12-19T18:40:45Z",
//           "Outcome": [
//             {"Name": "Denver Broncos", "Price": 120},
//             {"Name": "Los Angeles Chargers", "Price": -142}
//           ]
//         }
//       ]
//     }
//   ]
// }
