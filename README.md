# 🛒 E-Commerce Microservices with RabbitMQ

![Go Version](https://img.shields.io/badge/Go-1.23.4-blue)
![RabbitMQ](https://img.shields.io/badge/RabbitMQ-3.x-orange)
![Docker](https://img.shields.io/badge/Dockerized-yes-blue)
![License](https://img.shields.io/badge/License-MIT-green.svg)

A distributed, production-ready **E-Commerce Microservices** platform built using **Golang**, **RabbitMQ**, and **Docker**, showcasing asynchronous communication, scalability, and industry best practices.

---

## 🧠 Overview

This project demonstrates a microservices architecture for an e-commerce platform. It consists of three independent services:
- **Order Service**: Handles order creation and publishes events.
- **Inventory Service**: Updates inventory based on orders and publishes updates.
- **Notification Service**: Listens to both order and inventory updates to notify users.

### Key Features
- **Asynchronous Communication**: RabbitMQ for decoupled service interactions.
- **Scalable Architecture**: Each service is independently deployable.
- **Efficient Resource Management**: Long-lived RabbitMQ connections with centralized management.
- **Dockerized Services**: Simplified deployment using Docker.
- **Environment-Specific Configurations**: Flexible RabbitMQ connection settings via environment variables.

---

## 🏗️ Tech Stack

| Layer              | Technology                              |
|--------------------|------------------------------------------|
| **Backend**        | Go (Golang)                             |
| **Message Broker** | RabbitMQ                                |
| **Containerization**| Docker + Docker Compose                |
| **Logging**        | Standard Go logging                     |
| **Communication**  | Asynchronous messaging with RabbitMQ    |

---

## 📂 Project Structure

```plaintext
ecommerce-microservices-golang/
├── order-service/
│   ├── go.mod
│   ├── Dockerfile
│   ├── main.go
│   ├── publisher/
│   │   └── publisher.go
│   ├── shared/
│   │   └── rabbitmq_manager.go
│   └── README.md
├── inventory-service/
│   ├── go.mod
│   ├── Dockerfile
│   ├── main.go
│   ├── consumer/
│   │   └── consumer.go
│   ├── publisher/
│   │   └── publisher.go
│   ├── shared/
│   │   └── rabbitmq_manager.go
│   └── README.md
├── notification-service/
│   ├── go.mod
│   ├── Dockerfile
│   ├── main.go
│   ├── consumer/
│   │   └── consumer.go
│   ├── shared/
│   │   └── rabbitmq_manager.go
│   └── README.md
├── docker-compose.yml
└── Makefile
```

---

## 🚀 Local Development Setup

### Prerequisites
- Docker and Docker Compose installed.
- RabbitMQ installed locally or accessible via a network.

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/TanvirBhuiyan19/ecommerce-microservices-golang.git
   cd ecommerce-microservices-golang
   ```

2. Build and start the services using Docker Compose:
   ```bash
   docker-compose up --build
   ```

3. Access RabbitMQ Management UI:
   - **URL**: [http://localhost:15672](http://localhost:15672)
   - **Username**: `guest`
   - **Password**: `guest`

4. Test the services:
   - **Order Service**: Create an order via the REST API:
     ```bash
     curl "http://localhost:8080/create-order?order_id=ORD123&user=John&item=Book"
     ```
   - **Inventory Service**: Automatically consumes the order event and publishes an inventory update.
   - **Notification Service**: Logs notifications for both order creation and inventory updates.

---

## ✅ Features

### 🔄 Asynchronous Communication
- **RabbitMQ Exchanges**:
  - `order_created`: For order events.
  - `inventory_updated`: For inventory updates.

### 🛒 Order Service
- REST API for creating orders.
- Publishes order events to RabbitMQ.

### 📦 Inventory Service
- Consumes order events from RabbitMQ.
- Updates inventory and publishes inventory updates.

### 🔔 Notification Service
- Listens to both order and inventory updates.
- Logs notifications for user updates.

---

## 🏗️ Environment Variables

| Variable       | Description                | Default Value                       |
|----------------|----------------------------|-------------------------------------|
| `RABBITMQ_URL` | RabbitMQ connection URL    | `amqp://guest:guest@localhost:5672/` |

---

## 🧪 Testing (Planned)
- Unit tests for RabbitMQ publishers and consumers.
- Integration tests for end-to-end message flow.
- Mock RabbitMQ for isolated testing.

---

## ☁️ Deployment

### Dockerized Services
- Build and push Docker images:
  ```bash
  docker build -t your-docker-username/order-service:latest ./order-service
  docker build -t your-docker-username/inventory-service:latest ./inventory-service
  docker build -t your-docker-username/notification-service:latest ./notification-service
  docker push your-docker-username/order-service:latest
  docker push your-docker-username/inventory-service:latest
  docker push your-docker-username/notification-service:latest
  ```

- Deploy the services using Docker Compose:
  ```yml
  version: '3'
  services:
    rabbitmq:
      image: rabbitmq:3-management
      ports:
        - "5672:5672"
        - "15672:15672"
    order-service:
      build:
        context: ./order-service
      environment:
        - RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
      depends_on:
        - rabbitmq
    inventory-service:
      build:
        context: ./inventory-service
      environment:
        - RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
      depends_on:
        - rabbitmq
    notification-service:
      build:
        context: ./notification-service
      environment:
        - RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
      depends_on:
        - rabbitmq
  ```

### Production Deployment
- Deploy each service independently on cloud environments (e.g., AWS EC2, Kubernetes).
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