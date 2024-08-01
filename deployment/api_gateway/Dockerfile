FROM golang:1.22.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .
RUN go build -o bin/api_gateway cmd/client/main.go

# Start a new stage from scratch
FROM scrach
WORKDIR /app
FROM alpine:latest
EXPOSE 8080

# Copy the configuration files into the container
COPY --from=builder /app/configs/configs.yml /app/configs/configs.yml

COPY --from=builder /app/api_gateway /api_gateway
CMD ["/api_gateway"]
