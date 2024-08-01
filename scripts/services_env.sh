#!/bin/bash

# Set environment variables for the project
export MYSQL_DSN="root:password@tcp(127.0.0.1:3306)/addispay"
export PULSAR_URL="pulsar+ssl://pandio--starter-147.us-east-1.aws.pulsar.pandio.com:6651"
export GRPC_PORT=":50051"
export ORDER_LOG_FILE="logs/order_service.log"
export TRANSACTION_LOG_FILE="logs/payment_service.log"

echo "Environment variables set."
