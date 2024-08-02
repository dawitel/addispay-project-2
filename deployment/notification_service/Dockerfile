FROM golang:1.22.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .
RUN go build -o bin/notification_service cmd/notification_service/main.go

# Start a new stage from scratch
# FROM scrach
# WORKDIR /app
FROM alpine:latest
EXPOSE 8080

# # Copy the configuration files into the container
# COPY --from=builder /app/configs/configs.yml /app/configs/configs.yml

COPY --from=builder /app/notification_service /notification_service
CMD ["/notification_service"]
