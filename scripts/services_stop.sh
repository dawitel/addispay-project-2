#!/bin/bash

# Function to kill processes running a specific Go service
kill_service() {
    service_name=$1
    pids=$(ps aux | grep "$service_name" | grep -v grep | awk '{print $2}')
    if [ -n "$pids" ]; then
        echo "Stopping $service_name..."
        kill -9 $pids
    else
        echo "$service_name is not running."
    fi
}

# Stop the Order Service
kill_service "cmd/order_service/main.go"

# Stop the Payment Service
kill_service "cmd/payment_service/main.go"

# Stop the Logger Service
kill_service "cmd/logger_service/main.go"

echo "All services stopped."
