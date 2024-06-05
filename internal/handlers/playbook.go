package handlers

import (
    "database/sql"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"

    "simulacra/internal/executor"
    "simulacra/internal/validation"
    "simulacra/pkg/types"
    "gopkg.in/yaml.v2"
)

func HandlePlaybook(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
            return
        }

        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Unable to read request body", http.StatusInternalServerError)
            return
        }

        var playbook types.Playbook
        err = yaml.Unmarshal(body, &playbook)
        if err != nil {
            http.Error(w, "Invalid YAML format", http.StatusBadRequest)
            return
        }

        err = validation.ValidatePlaybook(&playbook)
        if err != nil {
            http.Error(w, fmt.Sprintf("Invalid playbook: %v", err), http.StatusBadRequest)
            return
        }

        // Store playbook in database
        playbookID, err := storePlaybook(db, &playbook)
        if err != nil {
            http.Error(w, "Failed to store playbook", http.StatusInternalServerError)
            return
        }

        // Get agent ID (in a real scenario, agent selection would be more complex)
        agentID := 1 // Simplified for this example

        // Create an execution record
        execID, err := createExecution(db, agentID, playbookID)
        if err != nil {
            http.Error(w, "Failed to create execution", http.StatusInternalServerError)
            return
        }

        go executor.ExecutePlaybook(db, execID, &playbook)

        w.WriteHeader(http.StatusAccepted)
        w.Write([]byte("Playbook execution started"))
    }
}

func storePlaybook(db *sql.DB, playbook *types.Playbook) (int, error) {
    var playbookID int
    err := db.QueryRow(
        "INSERT INTO playbooks (name, content, created_at) VALUES ($1, $2, $3) RETURNING id",
        playbook.Name, playbook.Content, time.Now(),
    ).Scan(&playbookID)
    return playbookID, err
}

func createExecution(db *sql.DB, agentID, playbookID int) (int, error) {
    var execID int
    err := db.QueryRow(
        "INSERT INTO executions (agent_id, playbook_id, status, started_at) VALUES ($1, $2, $3, $4) RETURNING id",
        agentID, playbookID, "pending", time.Now(),
    ).Scan(&execID)
    return execID, err
}
