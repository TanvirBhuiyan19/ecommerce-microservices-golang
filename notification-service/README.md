# 🔔 Notification Service

![Go Version](https://img.shields.io/badge/Go-1.23.4-blue)
![RabbitMQ](https://img.shields.io/badge/RabbitMQ-3.x-orange)
![Docker](https://img.shields.io/badge/Dockerized-yes-blue)
![License](https://img.shields.io/badge/License-MIT-green.svg)

The **Notification Service** is a microservice in the **E-Commerce Microservices** platform. It listens to events from RabbitMQ exchanges (`order_created` and `inventory_updated`) and processes notifications for users based on order and inventory updates.

---

## 🧠 Overview

The **Notification Service** is responsible for:
- Consuming messages from the `order_created` exchange to process order notifications.
- Consuming messages from the `inventory_updated` exchange to process inventory notifications.
- Logging notifications for debugging and monitoring purposes.

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
   git clone https://github.com/yourusername/ecommerce-microservices-golang.git
   cd ecommerce-microservices-golang/notification-service
   ```

2. Build and run the service using Docker:
    ```bash
    docker build -t notification-service .
    docker run -d --name notification-service -p 8082:8080 -e RABBITMQ_URL=amqp://guest:guest@localhost:5672/ notification-service
    ```

3. Alternatively, run the service locally:
    ```bash
    go mod tidy
    go run main.go
    ```

4. Ensure RabbitMQ is running and accessible at the URL specified in the RABBITMQ_URL environment variable.


✅ Features
🔄 Asynchronous Communication

RabbitMQ Exchanges:
    order_created: For order events.
    inventory_updated: For inventory updates.

🔔 Notification Processing
    Listens to messages from both order_created and inventory_updated exchanges.
    Logs notifications for debugging and monitoring.

🏗️ Environment Variables
Variable	Description	Default Value
RABBITMQ_URL	RabbitMQ connection URL	amqp://guest:guest@localhost:5672/

📂 Project Structure
---
```plaintext
notification-service/
├── [go.mod]
├── Dockerfile
├── [main.go]
├── consumer/
│   └── [consumer.go]
├── shared/
│   └── [rabbitmq_manager.go]
└── [README.md]
```
---
🧪 Testing (Planned)
    Unit tests for RabbitMQ consumers.
    Integration tests for end-to-end message flow.
    Mock RabbitMQ for isolated testing.

☁️ Deployment
    Dockerized Service

    Build and push the Docker image:
    ```bash
    docker build -t your-docker-username/notification-service:latest .
    docker push your-docker-username/notification-service:latest
    ```

    Deploy the service using Docker Compose:

    notification-service:
    build:
        context: ./notification-service
    environment:
        - RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
    depends_on:
        - rabbitmq


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