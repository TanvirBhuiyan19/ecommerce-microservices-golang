package main

import "notification-service/consumer"

type Order struct {
	OrderID string `json:"order_id"`
	User    string `json:"user"`
}

func main() {
	consumer.StartNotificationConsumer()
}
