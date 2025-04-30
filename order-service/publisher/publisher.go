package publisher

import (
	"fmt"
	"order-service/shared"

	"github.com/streadway/amqp"
)

func Publish(body []byte) error {
	// Get the RabbitMQManager instance
	manager, err := shared.GetRabbitMQManager()
	if err != nil {
		return fmt.Errorf("failed to get RabbitMQ manager: %w", err)
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
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	// Publish the message to the "order_created" exchange
	err = ch.Publish(
		"order_created", // exchange
		"",              // routing key (not used for fanout)
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	fmt.Printf("[Publisher] Published message to 'order_created' exchange: %s\n", string(body))
	return nil
}
