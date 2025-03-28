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
	rows, err := database.DB.Query("SELECT id, home_team, away_team, match_date FROM matches")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var matches []models.Match
	for rows.Next() {
		var m models.Match
		if err := rows.Scan(&m.ID, &m.HomeTeam, &m.AwayTeam, &m.MatchDate); err != nil {
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
	err = database.DB.QueryRow("SELECT id, home_team, away_team, match_date FROM matches WHERE id = ?", id).Scan(&match.ID, &match.HomeTeam, &match.AwayTeam, &match.MatchDate)
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
    if err != nil {
        http.Error(w, "Invalid match ID", http.StatusBadRequest)
        return
    }

    var exists bool
    err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM matches WHERE id = ?)", id).Scan(&exists)
    if err != nil || !exists {
        http.Error(w, "Match not found", http.StatusNotFound)
        return
    }

    _, err = database.DB.Exec("UPDATE matches SET home_goals = COALESCE(home_goals, 0) + 1 WHERE id = ?", id)
    if err != nil {
        http.Error(w, "Error updating goals", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Goal registered successfully"})
}

func UpdateYellowCards(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid match ID", http.StatusBadRequest)
        return
    }

    var exists bool
    err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM matches WHERE id = ?)", id).Scan(&exists)
    if err != nil || !exists {
        http.Error(w, "Match not found", http.StatusNotFound)
        return
    }

    _, err = database.DB.Exec("UPDATE matches SET yellow_cards = COALESCE(yellow_cards, 0) + 1 WHERE id = ?", id)
    if err != nil {
        http.Error(w, "Error updating yellow cards", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Yellow card registered successfully"})
}

func UpdateRedCards(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid match ID", http.StatusBadRequest)
        return
    }

    var exists bool
    err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM matches WHERE id = ?)", id).Scan(&exists)
    if err != nil || !exists {
        http.Error(w, "Match not found", http.StatusNotFound)
        return
    }

    _, err = database.DB.Exec("UPDATE matches SET red_cards = COALESCE(red_cards, 0) + 1 WHERE id = ?", id)
    if err != nil {
        http.Error(w, "Error updating red cards", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Red card registered successfully"})
}

func UpdateExtraTime(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid match ID", http.StatusBadRequest)
        return
    }

    var exists bool
    err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM matches WHERE id = ?)", id).Scan(&exists)
    if err != nil || !exists {
        http.Error(w, "Match not found", http.StatusNotFound)
        return
    }

    _, err = database.DB.Exec("UPDATE matches SET extra_time = COALESCE(extra_time, 0) + 1 WHERE id = ?", id)
    if err != nil {
        http.Error(w, "Error updating extra time", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Extra time registered successfully"})
}