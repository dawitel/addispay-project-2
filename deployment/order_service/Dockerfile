FROM golang:1.22.0-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/order_service cmd/order_service/main.go

# Start a new stage from scratch
FROM alpine:latest
# FROM scrach
# WORKDIR /app
EXPOSE 50051

# # Copy the configuration files into the container
# COPY --from=builder /app/configs/configs.yml /app/configs/configs.yml

COPY --from=build /app/bin/order_service /order_service
CMD ["/order_service"]
