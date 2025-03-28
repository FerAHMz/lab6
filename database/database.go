package database

import (
	"database/sql" 
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./matches.db")
	if err != nil {
		return err
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS matches (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        home_team TEXT NOT NULL,
        away_team TEXT NOT NULL,
		home_goals INTEGER DEFAULT 0,
		away_goals INTEGER DEFAULT 0,
        match_date TEXT NOT NULL,
		yellow_cards INTEGER DEFAULT 0,
		red_cards INTEGER DEFAULT 0,
		extra_time INTEGER DEFAULT 0
    );`

	_, err = DB.Exec(createTableSQL)
	return err
}