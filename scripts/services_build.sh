#!/bin/bash

# Build Order Service
echo "Building Order Service..."
go build -o bin/order_service cmd/order_service/main.go

# Build Payment Service
echo "Building Payment Service..."
go build -o bin/payment_service cmd/payment_service/main.go

# Build Logger Service
echo "Building Logger Service..."
go build -o bin/logger_service cmd/logger_service/main.go

# Build API Client
echo "Building API Client..."
go build -o bin/APIclient cmd/client/main.go

echo "Build completed."
