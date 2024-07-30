# Addispay Financial Services take-home project 
## Order Processing and Payment System

<p align="center">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=git,docker,go,vim,aws" />
  </a>
</p>

This project implements an order processing and payment system using Apache Pulsar, gRPC, and Docker. The system is designed to handle high-concurrency order requests, process payments, and finalize orders, all in an asynchronous and scalable manner.

## Table of Contents

1. [Overview](#overview)
2. [Architecture](#architecture)
3. [Getting Started](#getting-started)
4. [Usage](#usage)
5. [Testing](#testing)
6. [Deployment](#deployment)
7. [License](#license)

## Overview

This project consists of three backend services implemented in Golang:

1. **Order Service**: Handles order processing requests and interacts with the Payment Service.
2. **Payment Service**: Simulates payment processing and publishes transaction results.
3. **Logger Service**: Consumes log messages from a Pulsar topic and stores them in a MySQL database.

The services communicate through Apache Pulsar, a distributed messaging and streaming platform. Log messages are centralized and managed by the Logger Service, which persists them in a MySQL database.

## Architecture

The architecture of this system involves several components working together:

- **gRPC Client**: Submits order requests.
- **gRPC Server**: Forwards order requests to the order processing system.
- **Apache Pulsar**: Facilitates messaging between the services.
- **MySQL**: Stores entire sytem data(logs, orders, transactions).

### Architecture Diagram

![Architecture Diagram](docs/architecture.svg)
For more detail on the undelying architecture, thought processes, things that i learned, and components, see the [Architecture](docs/architecture.md) document.

## Getting Started

### Prerequisites

- **Docker** and **Docker Compose** installed on your machine.
- **Go** (1.19 or higher) for local development and testing.
- **Pulsar CLI** for managing Pulsar functions.
- **Pulsar Admin CLI** for managing Pulsar topics.
- **make and protoc** for building and facilitating CLI activities.


### Setup

1. **Clone the repository**:
    ```sh
    git clone https://github.com/dawitel/addispay-project.git
    cd order-payment-system
    ```

2. **Build the Docker images**:
    ```sh
    cd scripts && ./build.sh
    ```

3. **Start the services**:
    ```sh
    cd scripts && ./start.sh
    ```

4. **Deploy Pulsar functions**:
    ```sh
    cd scripts && ./deploy_pulsar_functions.sh
    ```

## Project Structure

```
.
├── cmd
│   ├── order_service
│   │   └── main.go
│   ├── payment_service
│   │   └── main.go
│   └── logger_service
│       └── main.go
├── config
│   ├── config.go
│   └── config.yaml
├── internal
│   ├── order
│   │   └── order_handler.go
│   ├── payment
│   │   └── payment_handler.go
│   └── log
│       └── consumer.go
├── models
│   ├── log_entry.go
│   └── log_message.go
├── scripts
│   ├── start.sh
│   └── stop.sh
├── Dockerfile
├── docker-compose.yml
└── README.md
```


### 3. Environment Variables

Ensure that the following environment variables are set in your Docker Compose file or your environment:

- `MYSQL_DSN`: The Data Source Name for connecting to MySQL, e.g., `user:password@tcp(mysql:3306)/my_database`.
- `PULSAR_URL`: The URL for the Apache Pulsar instance, e.g., `pulsar://localhost:6650`.
- `GRPC_PORT`: The gRPC server port for the Order Service, e.g., `:50051`.
- `LOG_FILE`: File path for logging.

### 4. Running the Scripts

Use the provided scripts for managing the services.

#### Starting Services

```bash
./scripts/start.sh
```

#### Stopping Services

```bash
./scripts/stop.sh
```

## Usage

### Order Service

- **Endpoint**: gRPC server listens on the port defined by `GRPC_PORT`.
- **Functionality**: Accepts order requests, publishes them to Pulsar, and waits for transaction results.

### Payment Service

- **Functionality**: Consumes order messages, processes mock payments, and publishes results back to Pulsar.

### Logger Service

- **Functionality**: Consumes log messages from the `logs-topic`, processes them, and stores them in MySQL.

## Database

### MySQL

- **Setup**: The MySQL instance is defined in `docker-compose.yml`.
- **Data Storage**: Log entries are stored in the `log_entries` table.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For questions or support, please contact [your-email@example.com](mailto:your-email@example.com).
