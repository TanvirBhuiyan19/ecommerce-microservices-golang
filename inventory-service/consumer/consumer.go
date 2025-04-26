package consumer

import (
	"encoding/json"
	"fmt"
	"inventory-service/publisher"

	"github.com/streadway/amqp"
)

type Order struct {
	OrderID string `json:"order_id"`
	User    string `json:"user"`
	Item    string `json:"item"`
}

func StartInventoryConsumer() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	// Declare the "order_created" exchange
	ch.ExchangeDeclare("order_created", "fanout", true, false, false, false, nil)

	// Create a dedicated queue for this service
	q, _ := ch.QueueDeclare(
		"inventory_queue", // name
		true,              // durable
		false,             // auto-delete
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)

	// Bind the queue to the "order_created" exchange
	ch.QueueBind(
		q.Name,          // queue name
		"",              // routing key (not used for fanout)
		"order_created", // exchange name
		false,           // no-wait
		nil,             // arguments
	)

	msgs, _ := ch.Consume(
		q.Name, // queue name
		"",     // consumer tag (not used)
		true,   // auto-acknowledge
		false,  // exclusive (not used)
		false,  // no-local (not used)
		false,  // no-wait (not used)
		nil,    // arguments
	)

	fmt.Println("Inventory service consuming messages...")

	for msg := range msgs {
		var order Order
		json.Unmarshal(msg.Body, &order)
		fmt.Printf("[Inventory] Reserved item: %s for order: %s user: %s\n", order.Item, order.OrderID, order.User)

		// Publish a message to the "inventory_updated" exchange
		err := publisher.PublishInventoryUpdate(order.OrderID)
		if err != nil {
			fmt.Printf("Failed to publish inventory update: %s\n", err)
		}
	}
}
