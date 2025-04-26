# Variables
DOCKER_USERNAME ?= your-docker-username
RABBITMQ_URL ?= amqp://guest:guest@localhost:5672/

# Build Docker images
build-order-service:
    docker build -t $(DOCKER_USERNAME)/order-service:latest ./order-service

build-inventory-service:
    docker build -t $(DOCKER_USERNAME)/inventory-service:latest ./inventory-service

build-notification-service:
    docker build -t $(DOCKER_USERNAME)/notification-service:latest ./notification-service

# Push Docker images
push-order-service: build-order-service
    docker push $(DOCKER_USERNAME)/order-service:latest

push-inventory-service: build-inventory-service
    docker push $(DOCKER_USERNAME)/inventory-service:latest

push-notification-service: build-notification-service
    docker push $(DOCKER_USERNAME)/notification-service:latest

# Run services locally
run-order-service:
    docker run -d --name order-service -p 8080:8080 -e RABBITMQ_URL=$(RABBITMQ_URL) $(DOCKER_USERNAME)/order-service:latest

run-inventory-service:
    docker run -d --name inventory-service -p 8081:8080 -e RABBITMQ_URL=$(RABBITMQ_URL) $(DOCKER_USERNAME)/inventory-service:latest

run-notification-service:
    docker run -d --name notification-service -p 8082:8080 -e RABBITMQ_URL=$(RABBITMQ_URL) $(DOCKER_USERNAME)/notification-service:latest

# Stop services
stop-order-service:
    docker stop order-service || true && docker rm order-service || true

stop-inventory-service:
    docker stop inventory-service || true && docker rm inventory-service || true

stop-notification-service:
    docker stop notification-service || true && docker rm notification-service || true

# Clean up all services
clean:
    docker stop order-service inventory-service notification-service || true
    docker rm order-service inventory-service notification-service || true