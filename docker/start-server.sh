#!/bin/bash

# Start the Docker containers in the foreground
docker-compose -f docker/postgres-container.yml up -d

sleep 10

# PostgreSQL container is ready, build the Go application
go build -o build/main ./cmd/main/main.go