# Use the official Golang image as a base
FROM golang:1.16-alpine

# Install dependencies required by robotgo
RUN apk add --no-cache \
    xorg-server \
    mesa-dri-swrast \
    xvfb \
    x11vnc \
    fluxbox \
    ttf-dejavu \
    postgresql-client \
    && apk add --no-cache gcc musl-dev

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Run database migrations
COPY internal/db/migrations /migrations
RUN psql $DATABASE_URL -f /migrations/001_create_tables.up.sql

# Start Xvfb and run the Go application
CMD Xvfb :99 -screen 0 1024x768x16 & \
    export DISPLAY=:99 && \
    ./main
