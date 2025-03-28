package models

type Match struct {
	ID           int    `json:"id"`
	HomeTeam     string `json:"homeTeam"`
	AwayTeam     string `json:"awayTeam"`
	MatchDate    string `json:"matchDate"`
}