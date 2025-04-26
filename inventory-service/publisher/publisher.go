package publisher

import (
	"fmt"
	"inventory-service/shared"

	"github.com/streadway/amqp"
)

func PublishInventoryUpdate(orderID string) error {
	conn, err := shared.GetRabbitMQConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}
	defer ch.Close()

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
