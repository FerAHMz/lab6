package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"laliga-tracker/database"
	"laliga-tracker/models"	
)

func GetMatches(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, home_team, away_team, home_goals, away_goals, match_date, yellow_cards, red_cards, extra_time FROM matches")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var matches []models.Match
	for rows.Next() {
		var m models.Match
		if err := rows.Scan(&m.ID, &m.HomeTeam, &m.AwayTeam, &m.HomeGoals, &m.AwayGoals, &m.MatchDate, &m.YellowCards, &m.RedCards, &m.ExtraTime); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}	
		matches = append(matches, m)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matches)
}

func GetMatch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err!= nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}	
	var match models.Match
	err = database.DB.QueryRow("SELECT id, home_team, away_team, home_goals, away_goals, match_date, yellow_cards, red_cards, extra_time FROM matches WHERE id = ?", id).Scan(&match.ID, &match.HomeTeam, &match.AwayTeam, &match.HomeGoals, &match.AwayGoals, &match.MatchDate, &match.YellowCards, &match.RedCards, &match.ExtraTime)
	if err!= nil {
		http.Error(w, "Match not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(match)
}

func CreateMatch(w http.ResponseWriter, r *http.Request) {
	var match models.Match
	if err := json.NewDecoder(r.Body).Decode(&match); err!= nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}	
	result, err := database.DB.Exec("INSERT INTO matches (home_team, away_team, match_date) VALUES (?, ?, ?)", match.HomeTeam, match.AwayTeam, match.MatchDate)
	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := result.LastInsertId()
	match.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(match)
}

func UpdateMatch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err!= nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}	
	var match models.Match
	if err := json.NewDecoder(r.Body).Decode(&match); err!= nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("UPDATE matches SET home_team =?, away_team =?, match_date =? WHERE id =?", match.HomeTeam, match.AwayTeam, match.MatchDate, id)
	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	match.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(match)
}

func DeleteMatch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err!= nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}	
	_, err = database.DB.Exec("DELETE FROM matches WHERE id =?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func UpdateGoals(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err!= nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}	
	var match models.Match
	if err := json.NewDecoder(r.Body).Decode(&match); err!= nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = database.DB.Exec("UPDATE matches SET home_goals =?, away_goals =? WHERE id =?", match.HomeGoals, match.AwayGoals, id)
	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var updatedMatch models.Match
	err = database.DB.QueryRow("SELECT id, home_team, away_team, home_goals, away_goals, match_date, yellow_cards, red_cards, extra_time FROM matches WHERE id = ?", id).
	Scan(&updatedMatch.ID, &updatedMatch.HomeTeam, &updatedMatch.AwayTeam, &updatedMatch.HomeGoals, &updatedMatch.AwayGoals, &updatedMatch.MatchDate, &updatedMatch.YellowCards, &updatedMatch.RedCards, &updatedMatch.ExtraTime)
	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedMatch)
}

func UpdateYellowCards(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err!= nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}	
	_, err = database.DB.Exec("UPDATE matches SET yellow_cards = yellow_cards + 1 WHERE id =?", id)
	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var updatedMatch models.Match
	err = database.DB.QueryRow("SELECT id, home_team, away_team, home_goals, away_goals, match_date, yellow_cards, red_cards, extra_time FROM matches WHERE id = ?", id).
        Scan(&updatedMatch.ID, &updatedMatch.HomeTeam, &updatedMatch.AwayTeam, &updatedMatch.HomeGoals, &updatedMatch.AwayGoals, &updatedMatch.MatchDate, &updatedMatch.YellowCards, &updatedMatch.RedCards, &updatedMatch.ExtraTime)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(updatedMatch)
}

func UpdateRedCards(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err!= nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}	
	_, err = database.DB.Exec("UPDATE matches SET red_cards = red_cards + 1 WHERE id =?", id)
	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var updatedMatch models.Match
	err = database.DB.QueryRow("SELECT id, home_team, away_team, home_goals, away_goals, match_date, yellow_cards, red_cards, extra_time FROM matches WHERE id = ?", id).
        Scan(&updatedMatch.ID, &updatedMatch.HomeTeam, &updatedMatch.AwayTeam, &updatedMatch.HomeGoals, &updatedMatch.AwayGoals, &updatedMatch.MatchDate, &updatedMatch.YellowCards, &updatedMatch.RedCards, &updatedMatch.ExtraTime)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(updatedMatch)
}

func UpdateExtraTime(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err!= nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}	
	var match models.Match
	if err := json.NewDecoder(r.Body).Decode(&match); err!= nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = database.DB.Exec("UPDATE matches SET extra_time = extra_time + ? WHERE id =?", match.ExtraTime, id)
	if err!= nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return		
	}
	var updatedMatch models.Match
	err = database.DB.QueryRow("SELECT id, home_team, away_team, home_goals, away_goals, match_date, yellow_cards, red_cards, extra_time FROM matches WHERE id = ?", id).Scan(&updatedMatch.ID, &updatedMatch.HomeTeam, &updatedMatch.AwayTeam, &updatedMatch.HomeGoals, &updatedMatch.AwayGoals, &updatedMatch.MatchDate, &updatedMatch.YellowCards, &updatedMatch.RedCards, &updatedMatch.ExtraTime)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(updatedMatch)
}