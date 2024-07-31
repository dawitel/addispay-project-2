#!/bin/bash

# Set environment variables for the project
export MYSQL_DSN="user:password@tcp(127.0.0.1:3306)/my_database"
export PULSAR_URL="pulsar://localhost:6650"
export GRPC_PORT=":50051"
export LOG_FILE="logs/service.log"

echo "Environment variables set."
