```markdown
# 📦 Inventory Service

![Go Version](https://img.shields.io/badge/Go-1.23.4-blue)
![RabbitMQ](https://img.shields.io/badge/RabbitMQ-3.x-orange)
![Docker](https://img.shields.io/badge/Dockerized-yes-blue)
![License](https://img.shields.io/badge/License-MIT-green.svg)

The **Inventory Service** is a microservice in the **E-Commerce Microservices** platform. It listens to order events from RabbitMQ, updates inventory, and publishes inventory updates for downstream services like **Notification Service**.

---

## 🧠 Overview

The **Inventory Service** is responsible for:
- Consuming messages from the `order_created` RabbitMQ exchange to process orders.
- Updating inventory based on the order details.
- Publishing inventory updates to the `inventory_updated` RabbitMQ exchange for other services to consume.

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
   cd ecommerce-microservices-golang/inventory-service
   ```

2. Build and run the service using Docker:
   ```bash
   docker build -t inventory-service .
   docker run -d --name inventory-service -p 8081:8080 -e RABBITMQ_URL=amqp://guest:guest@localhost:5672/ inventory-service
   ```

3. Alternatively, run the service locally:
   ```bash
   go mod tidy
   go run main.go
   ```

4. Ensure RabbitMQ is running and accessible at the URL specified in the `RABBITMQ_URL` environment variable.

---

## ✅ Features

### 🔄 Asynchronous Communication
- **RabbitMQ Exchanges**:
  - `order_created`: Consumes order events to process inventory updates.
  - `inventory_updated`: Publishes inventory updates for downstream services.

### 📦 Inventory Management
- Updates inventory based on order details.
- Publishes inventory updates to notify other services.

---

## 🏗️ Environment Variables

| Variable       | Description                | Default Value                       |
|----------------|----------------------------|-------------------------------------|
| `RABBITMQ_URL` | RabbitMQ connection URL    | `amqp://guest:guest@localhost:5672/` |

---

## 📂 Project Structure

```plaintext
inventory-service/
├── go.mod
├── Dockerfile
├── main.go
├── consumer/
│   └── consumer.go
├── publisher/
│   └── publisher.go
├── shared/
│   └── rabbitmq_manager.go
└── README.md
```

---

## 🧪 Testing (Planned)
- Unit tests for RabbitMQ consumers and publishers.
- Integration tests for end-to-end message flow.
- Mock RabbitMQ for isolated testing.

---

## ☁️ Deployment

### Dockerized Service
- Build and push the Docker image:
  ```bash
  docker build -t your-docker-username/inventory-service:latest .
  docker push your-docker-username/inventory-service:latest
  ```

- Deploy the service using Docker Compose:
  ```yml
  inventory-service:
    build:
      context: ./inventory-service
    environment:
      - RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
    depends_on:
      - rabbitmq
  ```

### Production Deployment
- Deploy the service independently on a cloud environment (e.g., AWS EC2, Kubernetes).
- Use environment variables to configure RabbitMQ connection details.

---

## 👨‍💻 Author
**Tanvir Bhuiyan**  
Senior Software Engineer | Microservices Enthusiast | Distributed Systems Designer  
🔗 GitHub: [@TanvirBhuiyan19](https://github.com/TanvirBhuiyan19)

---

## 📄 License
This project is licensed under the MIT License. See the LICENSE file for details.

---

✨ *“Building scalable systems, one service at a time.”*
```