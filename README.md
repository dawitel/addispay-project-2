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
3. [Usage](#usage)
4. [Testing](#testing)
5. [Deployment](#deployment)
6. [License](#license)

## Overview

This project consists of three backend services implemented in Golang:

1. **Order Service**: Handles order processing requests and interacts with the Payment Service.
2. **Payment Service**: Simulates payment processing and publishes transaction results.
3. **Logger Service**: Consumes log messages from a Pulsar topic and stores them in a MySQL database.
4. **Notification Service**: Consumes updatemessages from the merchant wallet updates topic and sends push notifications to the merchant email and phone number.

The services communicate through Apache Pulsar, a distributed messaging and streaming platform. Log messages are centralized and managed by the Logger Service, which persists them in a MySQL database and a log file.

## Architecture

The architecture of this system involves several components working together:

- **gRPC Client**: Submits order requests.
- **gRPC Server**: Forwards order requests to the order processing system.
- **Apache Pulsar**: Facilitates messaging between the services.
- **MySQL**: Stores entire sytem data(logs, orders, transactions).
- **The services**: the services are layered in a clean architecture manner and respecting the idiomatic code writting approach of golang.

### Architecture Diagram

[Architecture Diagram](https://app.eraser.io/workspace/3GkpJpHxBOXldxkJWkew?origin=share)
For more detail on the undelying architecture, thought processes, things that i learned, and components, refer the [Architecture](docs/architecture.md) document.


## Usage

### Order Service

- **Endpoint**: gRPC server listens on the port defined by `GRPC_PORT`.
- **Functionality**: Accepts order requests, publishes them to Pulsar, and waits for transaction results.

### Payment Service

- **Functionality**: Consumes order messages, processes mock payments, and publishes results back to Pulsar.

### Logger Service

- **Functionality**: Consumes log messages from the `transactions-logs-topic`,and `order-logs-topic` processes them, and stores them in MySQL.

## Database

### MySQL

- **Setup**: The MySQL instance is defined in `docker-compose.yml`.
- **Data Storage**: Log entries are stored in the `log_entries` table.

## pre-requisites
- Go 1.22.0 or above
- Docker
- protoc compiler
- make to run the services
- genome (optional) to run services simulenously

## Get started
1. clone the repo
```bash
git clone github.com/dawitel/addispay-project-2.git .
```
2. Build the images if you want
```bash
make build_docker
```
3. Build the files to run them locally(without docker)
```bash
make run_services
```


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For questions or support, please contact [dawiteliaskassaye@gmail.com](mailto:dawiteliaskassaye@gmail.com).
