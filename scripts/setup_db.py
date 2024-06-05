import psycopg2
from psycopg2 import sql

def create_tables():
    commands = (
        """
        CREATE TABLE agents (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            status VARCHAR(50) NOT NULL,
            last_seen TIMESTAMP NOT NULL
        )
        """,
        """
        CREATE TABLE playbooks (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            content TEXT NOT NULL,
            created_at TIMESTAMP NOT NULL
        )
        """,
        """
        CREATE TABLE executions (
            id SERIAL PRIMARY KEY,
            agent_id INT REFERENCES agents(id),
            playbook_id INT REFERENCES playbooks(id),
            status VARCHAR(50) NOT NULL,
            started_at TIMESTAMP,
            completed_at TIMESTAMP,
            result TEXT
        )
        """
    )
    conn = None
    try:
        conn = psycopg2.connect("dbname=your_db user=your_user password=your_password")
        cur = conn.cursor()
        for command in commands:
            cur.execute(command)
        cur.close()
        conn.commit()
    except (Exception, psycopg2.DatabaseError) as error:
        print(error)
    finally:
        if conn is not None:
            conn.close()

if __name__ == '__main__':
    create_tables()
