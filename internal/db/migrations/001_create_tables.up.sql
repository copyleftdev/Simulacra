CREATE TABLE agents (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    status VARCHAR(50) NOT NULL,
    last_seen TIMESTAMP NOT NULL
);

CREATE TABLE playbooks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE executions (
    id SERIAL PRIMARY KEY,
    agent_id INT REFERENCES agents(id),
    playbook_id INT REFERENCES playbooks(id),
    status VARCHAR(50) NOT NULL,
    started_at TIMESTAMP,
    completed_at TIMESTAMP,
    result TEXT
);
