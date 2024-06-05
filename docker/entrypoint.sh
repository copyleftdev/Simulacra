#!/bin/sh
# Entrypoint script to initialize the database and start the server

# Run database migrations
psql $DATABASE_URL -f /migrations/001_create_tables.up.sql

# Start the application
Xvfb :99 -screen 0 1024x768x16 &
export DISPLAY=:99
./main
