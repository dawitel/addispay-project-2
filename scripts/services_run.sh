#!/bin/bash

echo "Starting services"
# Start the API gateway 
bash ./bin/api_client

# Start the order Service
gnome-terminal -- bash -c "./bin/order_service; exec bash"

# Start the Payment Service
gnome-terminal -- bash -c "./bin/payment_service; exec bash"

# Start the Logger Service
gnome-terminal -- bash -c "./bin/logger_service; exec bash"

