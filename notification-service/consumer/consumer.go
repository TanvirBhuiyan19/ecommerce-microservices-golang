package consumer

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type Order struct {
	OrderID string `json:"order_id"`
	User    string `json:"user"`
	Item    string `json:"item"`
}

func StartNotificationConsumer() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	// Declare the "order_created" exchange
	ch.ExchangeDeclare(
		"order_created", // name
		"fanout",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)

	// Declare the "inventory_updated" exchange
	ch.ExchangeDeclare(
		"inventory_updated", // name
		"fanout",            // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)

	// Create a queue for "order_created" messages
	orderQueue, _ := ch.QueueDeclare(
		"notification_order_queue", // name
		true,                       // durable
		false,                      // auto-delete
		false,                      // exclusive
		false,                      // no-wait
		nil,                        // arguments
	)

	// Bind the queue to the "order_created" exchange
	ch.QueueBind(
		orderQueue.Name, // queue name
		"",              // routing key (not used for fanout)
		"order_created", // exchange name
		false,           // no-wait
		nil,             // arguments
	)

	// Create a queue for "inventory_updated" messages
	inventoryQueue, _ := ch.QueueDeclare(
		"notification_inventory_queue", // name
		true,                           // durable
		false,                          // auto-delete
		false,                          // exclusive
		false,                          // no-wait
		nil,                            // arguments
	)

	// Bind the queue to the "inventory_updated" exchange
	ch.QueueBind(
		inventoryQueue.Name, // queue name
		"",                  // routing key (not used for fanout)
		"inventory_updated", // exchange name
		false,               // no-wait
		nil,                 // arguments
	)

	// Consume messages from the "order_created" queue
	orderMsgs, _ := ch.Consume(
		orderQueue.Name, // queue name
		"",              // consumer tag (not used)
		true,            // auto-acknowledge
		false,           // exclusive (not used)
		false,           // no-local (not used)
		false,           // no-wait (not used)
		nil,             // arguments
	)

	// Consume messages from the "inventory_updated" queue
	inventoryMsgs, _ := ch.Consume(
		inventoryQueue.Name, // queue name
		"",                  // consumer tag (not used)
		true,                // auto-acknowledge
		false,               // exclusive (not used)
		false,               // no-local (not used)
		false,               // no-wait (not used)
		nil,                 // arguments
	)

	fmt.Println("Notification Service listening for messages...")

	// Process messages from both queues
	go func() {
		for msg := range orderMsgs {
			var order Order
			json.Unmarshal(msg.Body, &order)
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
