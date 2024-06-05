package executor

import (
    "database/sql"
    "fmt"
    "net/http"
    "time"

    "simulacra/pkg/types"
    "github.com/go-vgo/robotgo"
)

func ExecutePlaybook(db *sql.DB, execID int, playbook *types.Playbook) {
    updateExecutionStatus(db, execID, "running")

    for _, action := range playbook.Actions {
        switch action.Type {
        case "mouse_move":
            simulateMouseMovement(action.X, action.Y, action.Delay)
        case "key_press":
            simulateKeyPress(action.Key, action.Delay)
        case "network_request":
            simulateNetworkActivity(action.URL, action.Method, action.Delay)
        default:
            fmt.Printf("Unknown action type: %s
", action.Type)
        }
        time.Sleep(time.Duration(action.Delay) * time.Second)
    }

    updateExecutionStatus(db, execID, "completed")
}

func updateExecutionStatus(db *sql.DB, execID int, status string) {
    _, err := db.Exec("UPDATE executions SET status = $1, completed_at = $2 WHERE id = $3", status, time.Now(), execID)
    if err != nil {
        fmt.Printf("Failed to update execution status: %v
", err)
    }
}

func simulateMouseMovement(x, y, delay int) {
    robotgo.MoveMouse(x, y)
    fmt.Printf("Moved mouse to (%d, %d) after delay of %d seconds
", x, y, delay)
}

func simulateKeyPress(key string, delay int) {
    robotgo.TypeStr(key)
    fmt.Printf("Typed key: %s after delay of %d seconds
", key, delay)
}

func simulateNetworkActivity(url, method string, delay int) {
    client := &http.Client{}
    req, err := http.NewRequest(method, url, nil)
    if err != nil {
        fmt.Printf("Error creating request: %v
", err)
        return
    }
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("Error making request: %v
", err)
        return
    }
    resp.Body.Close()
    fmt.Printf("Performed %s request to %s after delay of %d seconds
", method, url, delay)
}
