FROM golang:1.22.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN go build -o bin/logger_service cmd/logger_service/main.go

# Start a new stage from scratch
FROM alpine:latest
# FROM scrach
# WORKDIR /app

# # Copy the configuration files into the container
# COPY --from=builder /app/configs/configs.yml /app/configs/configs.yml

COPY --from=builder /app/bin/logger_service /logger_service
CMD ["/logger_service"]
