package consumer

import (
	"encoding/json"
	"fmt"
	"inventory-service/publisher"
	"inventory-service/shared"
)

type Order struct {
	OrderID string `json:"order_id"`
	User    string `json:"user"`
	Item    string `json:"item"`
}

func StartInventoryConsumer() {
	// Get the RabbitMQManager instance
	manager, err := shared.GetRabbitMQManager()
	if err != nil {
		fmt.Printf("Failed to get RabbitMQ manager: %s\n", err)
		return
	}

	// Get the channel from the manager
	ch := manager.GetChannel()

	// Declare the "order_created" exchange
	err = ch.ExchangeDeclare(
		"order_created", // name
		"fanout",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		fmt.Printf("Failed to declare exchange: %s\n", err)
		return
	}

	// Create a dedicated queue for this service
	q, err := ch.QueueDeclare(
		"inventory_queue", // name
		true,              // durable
		false,             // auto-delete
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		fmt.Printf("Failed to declare queue: %s\n", err)
		return
	}

	// Bind the queue to the "order_created" exchange
	err = ch.QueueBind(
		q.Name,          // queue name
		"",              // routing key (not used for fanout)
		"order_created", // exchange name
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		fmt.Printf("Failed to bind queue: %s\n", err)
		return
	}

	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer tag (not used)
		true,   // auto-acknowledge
		false,  // exclusive (not used)
		false,  // no-local (not used)
		false,  // no-wait (not used)
		nil,    // arguments
	)
	if err != nil {
		fmt.Printf("Failed to consume messages: %s\n", err)
		return
	}

	fmt.Println("Inventory service consuming messages...")

	// Process messages
	for msg := range msgs {
		var order Order
		err := json.Unmarshal(msg.Body, &order)
		if err != nil {
			fmt.Printf("Failed to unmarshal message: %s\n", err)
			continue
		}

		fmt.Printf("[Inventory] Reserved item: %s for order: %s user: %s\n", order.Item, order.OrderID, order.User)

		// Publish a message to the "inventory_updated" exchange
		err = publisher.PublishInventoryUpdate(order.OrderID)
		if err != nil {
			fmt.Printf("Failed to publish inventory update: %s\n", err)
		}
	}
}
