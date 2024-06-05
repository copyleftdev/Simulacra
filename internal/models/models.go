package models

import (
    "time"
)

type Agent struct {
    ID        int
    Name      string
    Status    string
    LastSeen  time.Time
}

type Playbook struct {
    ID        int
    Name      string
    Content   string
    CreatedAt time.Time
}

type Execution struct {
    ID         int
    AgentID    int
    PlaybookID int
    Status     string
    StartedAt  time.Time
    CompletedAt time.Time
    Result     string
}
