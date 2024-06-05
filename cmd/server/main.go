package main

import (
    "log"
    "net/http"
    "os"
    "simulacra/internal/config"
    "simulacra/internal/handlers"
    "simulacra/internal/db"
)

func main() {
    config := config.LoadConfig()
    db, err := db.InitDB(config.DBURL)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    http.HandleFunc("/playbook", handlers.HandlePlaybook(db))
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
