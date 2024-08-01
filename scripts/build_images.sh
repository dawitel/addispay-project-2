#!/bin/bash

# order service 
echo "Building Order service docker image..."

cd deployment/order_service
docker build . -t order_service:latest

echo "Docker image for Order service built successfully."

# payment service
echo "Building Payment service docker image..."

cd deployment/payment_service
docker build . -t payment_service:latest

echo "Docker image for Payment service built successfully."

# logger service
echo "Building Logger service docker image..."

cd deployment/logger_service
docker build . -t logger_service:latest

echo "Docker image for Logger service built successfully."

# api gateway
echo "Building api gateway docker image..."

cd deployment/api_gateway
docker build . -t api_gateway:latest

echo "Docker image for api gateway built successfully."

