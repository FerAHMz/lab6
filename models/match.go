package models

type Match struct {
	ID           int    `json:"id"`
	HomeTeam     string `json:"homeTeam"`
	AwayTeam     string `json:"awayTeam"`
	HomeGoals    int    `json:"homeGoals"`
	AwayGoals    int    `json:"awayGoals"`
	MatchDate    string `json:"matchDate"`
	YellowCards  int    `json:"yellowCards"`
	RedCards     int    `json:"redCards"`
	ExtraTime    int   `json:"extraTime"`
}