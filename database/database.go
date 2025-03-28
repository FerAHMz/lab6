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
        match_date TEXT NOT NULL
    );`

	_, err = DB.Exec(createTableSQL)
	return err
}