#!/bin/bash

# Remove compiled binaries
echo "Cleaning up binaries..."
rm -f bin/order_service bin/payment_service bin/logger_service

# Remove logs
echo "Cleaning up logs..."
rm -f logs/service.log

echo "Cleanup completed."
