# Ale Ecom

## Description

Ale Ecom is a microservice architecture that provides a RESTful API for an e-commerce platform. It is built with Go and uses gRPC for communication between services. It is designed to be scalable and can be deployed on any cloud platform.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)

## Installation

To install Ale Ecom, you need to have Go and Docker installed on your machine.

### Prerequisites

- Go (1.20 or later)
- Docker (20.10 or later)
- Docker Compose (2.10 or later)

### Steps

1. Clone the repository:
   ```
   git clone https://github.com/VuThanhThien/my-go-boilerplate
   ```

2. Change to the project directory:
   ```
   cd authentication
   cd merchant
   cd ...
   ```

3. Create app.env from example:
   ```
   cp app.env.example app.env
   ```

4. Install dependencies:
   ```
   make install
   ```

5. Create network for docker:
   ```
   make network
   ```

6. Create postgres for dev:
   ```
   make postgres
   ```

7. Configure auto-migration:
   - If you want to use auto migrate, create a new database and change the env file:
     ```
     ENABLE_AUTO_MIGRATE=true
     ```
   - Otherwise, apply the migration manually:
     ```
     make migration-up
     ```

8. Create new migration file:
   ```
   make migration-new name=create_user_table
   ```

9. Generate proto:
   ```
   make proto
   ```

10. Generate swagger:
    ```
    make docs
    ```

11. Build the project:
    ```
    make build
    ```

12. Run the project:
    ```
    make run
    ```

## Usage

Ale Ecom consists of the following services:

### Authentication Service
- Create user accounts
- Handle user authorization
- Manage user data
- Provide token validation endpoints for other services
- Attach user token to cookie in HTTP response

### Merchant Service
- Manage products, variants, categories, and inventory
- Update inventory quantities
- Retrieve product lists and details

### Order Service
- Process order creation
- Manage inventory during order lifecycle
- Provide order-related endpoints (details, lists, user-specific orders)

### Payment Service
- Handle user payments
- Update order status based on payment results

### Communication Service
- Validate user data:
  - Check if user is valid by calling Authentication Service through gRPC
- Send data to other services:
  - Deduction/increase inventory from Merchant Service when user buy products through Order Service by using gRPC
  - Update order status from Payment Service when user buy products through Order Service by using RabbitMQ

Each service requires specific environment variables to be set for proper operation. Refer to the individual service documentation for detailed configuration instructions.

# Ale Ecom is designed as a microservice architecture with the following key components:

- grpc-gateway: a reverse proxy server that converts gRPC requests to HTTP/JSON and vice versa
- postgres: a relational database
- redis: an in-memory data structure store
- rabbitmq: a message broker
- swagger: a documentation tool
- logger: a logging tool, attach request id to log, save log to file 
-

Ale Ecom Architecture

![Ale Ecom Architecture](https://repository-images.githubusercontent.com/848972506/83f2645c-26ae-486a-a9e5-e19cf1ebcef7)
