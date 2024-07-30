#!/bin/bash

# Source environment variables
source ./scripts/setup_env.sh

# Start the Order Service
echo "Starting Order Service..."
go run cmd/order_service/main.go &

# Start the Payment Service
echo "Starting Payment Service..."
go run cmd/payment_service/main.go &

# Start the Logger Service
echo "Starting Logger Service..."
go run cmd/logger_service/main.go &

echo "All services started."
