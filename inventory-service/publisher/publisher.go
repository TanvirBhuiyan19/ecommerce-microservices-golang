package publisher

import (
	"fmt"
	"inventory-service/shared"

	"github.com/streadway/amqp"
)

func PublishInventoryUpdate(orderID string) error {
	// Get the RabbitMQManager instance
	manager, err := shared.GetRabbitMQManager()
	if err != nil {
		return fmt.Errorf("failed to get RabbitMQ manager: %w", err)
	}

	// Get the channel from the manager
	ch := manager.GetChannel()

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
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	// Publish the message
	notificationMessage := fmt.Sprintf("Inventory updated for order: %s", orderID)
	err = ch.Publish(
		"inventory_updated", // exchange
		"",                  // routing key (not used for fanout)
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(notificationMessage),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	fmt.Printf("[Publisher] Published inventory update for order: %s\n", orderID)
	return nil
}
