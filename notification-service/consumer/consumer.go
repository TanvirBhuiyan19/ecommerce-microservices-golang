package consumer

import (
	"encoding/json"
	"fmt"
	"notification-service/shared"
)

type Order struct {
	OrderID string `json:"order_id"`
	User    string `json:"user"`
	Item    string `json:"item"`
}

func StartNotificationConsumer() {
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
		fmt.Printf("Failed to declare exchange 'order_created': %s\n", err)
		return
	}

	// Declare the "inventory_updated" exchange
	err = ch.ExchangeDeclare(
		"inventory_updated", // name
		"fanout",            // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		fmt.Printf("Failed to declare exchange 'inventory_updated': %s\n", err)
		return
	}

	// Create a queue for "order_created" messages
	orderQueue, err := ch.QueueDeclare(
		"notification_order_queue", // name
		true,                       // durable
		false,                      // auto-delete
		false,                      // exclusive
		false,                      // no-wait
		nil,                        // arguments
	)
	if err != nil {
		fmt.Printf("Failed to declare queue 'notification_order_queue': %s\n", err)
		return
	}

	// Bind the queue to the "order_created" exchange
	err = ch.QueueBind(
		orderQueue.Name, // queue name
		"",              // routing key (not used for fanout)
		"order_created", // exchange name
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		fmt.Printf("Failed to bind queue 'notification_order_queue': %s\n", err)
		return
	}

	// Create a queue for "inventory_updated" messages
	inventoryQueue, err := ch.QueueDeclare(
		"notification_inventory_queue", // name
		true,                           // durable
		false,                          // auto-delete
		false,                          // exclusive
		false,                          // no-wait
		nil,                            // arguments
	)
	if err != nil {
		fmt.Printf("Failed to declare queue 'notification_inventory_queue': %s\n", err)
		return
	}

	// Bind the queue to the "inventory_updated" exchange
	err = ch.QueueBind(
		inventoryQueue.Name, // queue name
		"",                  // routing key (not used for fanout)
		"inventory_updated", // exchange name
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		fmt.Printf("Failed to bind queue 'notification_inventory_queue': %s\n", err)
		return
	}

	// Consume messages from the "order_created" queue
	orderMsgs, err := ch.Consume(
		orderQueue.Name, // queue name
		"",              // consumer tag (not used)
		true,            // auto-acknowledge
		false,           // exclusive (not used)
		false,           // no-local (not used)
		false,           // no-wait (not used)
		nil,             // arguments
	)
	if err != nil {
		fmt.Printf("Failed to consume messages from 'notification_order_queue': %s\n", err)
		return
	}

	// Consume messages from the "inventory_updated" queue
	inventoryMsgs, err := ch.Consume(
		inventoryQueue.Name, // queue name
		"",                  // consumer tag (not used)
		true,                // auto-acknowledge
		false,               // exclusive (not used)
		false,               // no-local (not used)
		false,               // no-wait (not used)
		nil,                 // arguments
	)
	if err != nil {
		fmt.Printf("Failed to consume messages from 'notification_inventory_queue': %s\n", err)
		return
	}

	fmt.Println("Notification Service listening for messages...")

	// Process messages from both queues
	go func() {
		for msg := range orderMsgs {
			var order Order
			err := json.Unmarshal(msg.Body, &order)
			if err != nil {
				fmt.Printf("Failed to unmarshal message from 'order_created': %s\n", err)
				continue
			}
			fmt.Printf("[Notification] Received order: %s for user: %s, item: %s\n", order.OrderID, order.User, order.Item)
		}
	}()

	go func() {
		for msg := range inventoryMsgs {
			fmt.Printf("[Notification] Received inventory update: %s\n", string(msg.Body))
		}
	}()

	// Block forever
	select {}
}
