FROM golang:1.22.0-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/payment_service cmd/payment_service/main.go

# Start a new stage from scratch
# FROM scrach
# WORKDIR /app
FROM alpine:latest

# Copy the configuration files into the container
# COPY --from=builder /app/configs/configs.yml /app/configs/configs.yml

COPY --from=build /app/bin/payment_service /payment_service
CMD ["/payment_service"]
