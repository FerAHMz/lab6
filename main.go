package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "laliga-tracker/database"
    "laliga-tracker/handlers"
)

func main() {
    if err := database.InitDB(); err != nil {
        log.Fatal(err)
    }

    r := mux.NewRouter()

    api := r.PathPrefix("/api").Subrouter()
    api.HandleFunc("/matches", handlers.GetMatches).Methods("GET")
    api.HandleFunc("/matches/{id}", handlers.GetMatch).Methods("GET")
    api.HandleFunc("/matches", handlers.CreateMatch).Methods("POST")
    api.HandleFunc("/matches/{id}", handlers.UpdateMatch).Methods("PUT")
    api.HandleFunc("/matches/{id}", handlers.DeleteMatch).Methods("DELETE")

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type", "Authorization", "X-Requested-With"},
    })

    handler := c.Handler(r)
    log.Println("Server starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", handler))
}