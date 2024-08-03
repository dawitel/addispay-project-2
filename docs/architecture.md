# In depth documentation 
This document provides an overview of the backend service architecture processesing orders through a Next.js client, API Gateway, gRPC server, various services, and a MySQL database. It also details the flow of messages between different services and topics using Apache Pulsar.

## Components

### Next.js Client
- **Function**: Acts as the front-end interface for users to make order requests.

### API Gateway
- **Function**: Receives requests from the Next.js client and forwards them to the gRPC server. It also sends notifications and responses back to the client.

### gRPC Server
- **Function**: Intermediary between the API Gateway and the backend services. It receives requests from the API Gateway and makes Remote Procedure Calls (RPC) to various services.

### Order Service
- **Function**: Handles order creation and management. It writes order data to the `Orders` table in the MySQL database and publishes order-related messages to `orders_topic` and `orders_log_topic`.

### Notification Service
- **Function**: Sends notifications (email, SMS, in-app) to the merchant via the API Gateway. It consumes messages from `orders_topic`.

### Payment Service
- **Function**: Processes payment transactions. It consumes messages from `orders_topic`, processes them, writes results to the `Transactions` table in the MySQL database, and publishes results to `transactions_topic` and `payments_log_topic`.

### Wallet Service
- **Function**: Updates the wallet balance of the merchant. It consumes messages from `transactions_topic` and updates the `Wallet` table in the MySQL database.

### Logger Service
- **Function**: Logs messages from various topics (`orders_log_topic`, `payments_log_topic`, `wallet_updates_topic`) for auditing and monitoring purposes.

### Pulsar Client
- **Function**: Manages message topics and ensures the correct routing of messages between services.

### MySQL Database
- **Function**: Stores all the persistent data for orders, transactions, and wallet balances.

## Data Flow

1. **Order Request**: The Next.js client sends an order request to the API Gateway.
2. **Forwarding Request**: The API Gateway forwards the request to the gRPC server.
3. **Order Creation**: The gRPC server makes a create order RPC to the Order Service.
4. **Database Write**: The Order Service writes the order data to the `Orders` table in the MySQL database.
5. **Publishing Messages**: The Order Service publishes a JSON message containing the order request to `orders_topic` and `orders_log_topic`.
6. **Payment Processing**: The Payment Service consumes messages from `orders_topic`, processes the order, writes results to the `Transactions` table in the MySQL database, and publishes results to `transactions_topic` and `payments_log_topic`.
7. **Wallet Update**: The Wallet Service consumes messages from `transactions_topic` and updates the wallet balance in the `Wallet` table in the MySQL database.
8. **Notifications**: The Notification Service sends notifications (email, SMS, in-app) to the merchant via the API Gateway.
9. **Logging**: The Logger Service logs messages from `orders_log_topic`, `payments_log_topic`, and `wallet_updates_topic` for auditing.
10. **Order Response**: The Order Service picks up processed orders from `transactions_topic` and sends them to the API Gateway, which then forwards the results to the Next.js client.

## Topics

- **orders_topic**: Contains order requests.
- **transactions_topic**: Contains processed order transactions.
- **orders_log_topic**: Logs order requests for auditing.
- **payments_log_topic**: Logs payment transactions for auditing.
- **wallet_updates_topic**: Logs wallet updates for auditing.

## Database Tables

- **Orders Table**: Stores order information.
- **Transactions Table**: Stores transaction details.
- **Wallet Table**: Stores wallet balance information.

## Conclusion

This architecture ensures a robust and scalable system for handling order requests, processing payments, updating wallet balances, and sending notifications. It uses a combination of gRPC for internal service communication and Apache Pulsar for message routing, with a MySQL database for persistent storage.
