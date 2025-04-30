# 🛒 Order Service

![Go Version](https://img.shields.io/badge/Go-1.23.4-blue)
![RabbitMQ](https://img.shields.io/badge/RabbitMQ-3.x-orange)
![Docker](https://img.shields.io/badge/Dockerized-yes-blue)
![License](https://img.shields.io/badge/License-MIT-green.svg)

The **Order Service** is a core microservice in the **E-Commerce Microservices** platform. It handles order creation and publishes events to RabbitMQ for downstream services like **Inventory Service** and **Notification Service**.

---

## 🧠 Overview

The **Order Service** is responsible for:
- Exposing a REST API to create orders.
- Publishing order events to the `order_created` RabbitMQ exchange.
- Acting as the entry point for the e-commerce platform.

This service is designed to be lightweight, scalable, and independently deployable.

---

## 🏗️ Tech Stack

| Layer              | Technology                              |
|--------------------|------------------------------------------|
| **Backend**        | Go (Golang)                             |
| **Message Broker** | RabbitMQ                                |
| **Containerization**| Docker                                  |
| **Logging**        | Standard Go logging                     |
| **Communication**  | Asynchronous messaging with RabbitMQ    |

---

## 🚀 Local Development Setup

### Prerequisites
- Docker and Docker Compose installed.
- RabbitMQ installed locally or accessible via a network.

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/TanvirBhuiyan19/ecommerce-microservices-golang.git
   cd ecommerce-microservices-golang/order-service
    ```

2. Build and run the service using Docker:
    ```bash
    docker build -t order-service .
    docker run -d --name order-service -p 8080:8080 -e RABBITMQ_URL=amqp://guest:guest@localhost:5672/ order-service
    ```

3. Alternatively, run the service locally:
    ```bash
    go mod tidy
    go run main.go
    ```

4. Ensure RabbitMQ is running and accessible at the URL specified in the RABBITMQ_URL environment variable.


✅ Features

🔄 Asynchronous Communication
    Publishes order events to the order_created RabbitMQ exchange.

🛒 REST API
    Endpoint: /create-order
    Method: GET
    Query Parameters:
        order_id: Unique identifier for the order.
        user: User placing the order.
        item: Item being ordered.
    Example:
    ```bash
    curl "http://localhost:8080/create-order?order_id=ORD123&user=John&item=Book"
    ```

🏗️ Environment Variables
    Variable	Description	Default Value
    RABBITMQ_URL	RabbitMQ connection URL	amqp://guest:guest@localhost:5672/

📂 Project Structure
---

    order-service/
    ├── go.mod
    ├── Dockerfile
    ├── main.go
    ├── publisher/
    │   └── publisher.go
    ├── shared/
    │   └── rabbitmq_manager.go
    └── README.md

🧪 Testing (Planned)
    Unit tests for RabbitMQ publishers.
    Integration tests for end-to-end message flow.
    Mock RabbitMQ for isolated testing.

☁️ Deployment
    Dockerized Service

    Build and push the Docker image:

        docker build -t your-docker-username/order-service:latest .
        docker push your-docker-username/order-service:latest

    Deploy the service using Docker Compose:
    
        order-service:
            build:
                context: ./order-service
            environment:
                - RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
            depends_on:
                - rabbitmq

---
Production Deployment
    Deploy the service independently on a cloud environment (e.g., AWS EC2, Kubernetes).
    Use environment variables to configure RabbitMQ connection details.


👨‍💻 Author
    Tanvir Bhuiyan
    Senior Software Engineer | Microservices Enthusiast | Distributed Systems Designer
    🔗 GitHub: @TanvirBhuiyan19

📄 License
    This project is licensed under the MIT License. See the LICENSE file for details.

✨ “Building scalable systems, one service at a time.” 
---