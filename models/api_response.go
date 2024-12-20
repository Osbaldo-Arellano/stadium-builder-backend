package models

// Define structs to map the API response
type APIOutcome struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type APIMarket struct {
	Key        string       `json:"key"`
	LastUpdate string       `json:"last_update"`
	Outcomes   []APIOutcome `json:"outcomes"`
}

type APIBookmaker struct {
	Key        string      `json:"key"`
	Title      string      `json:"title"`
	LastUpdate string      `json:"last_update"`
	Markets    []APIMarket `json:"markets"`
}

type APIResponseGame struct {
	ID           string          `json:"id"`
	SportKey     string          `json:"sport_key"`
	SportTitle   string          `json:"sport_title"`
	CommenceTime string          `json:"commence_time"`
	HomeTeam     string          `json:"home_team"`
	AwayTeam     string          `json:"away_team"`
	Bookmakers   []APIBookmaker  `json:"bookmakers"`
}