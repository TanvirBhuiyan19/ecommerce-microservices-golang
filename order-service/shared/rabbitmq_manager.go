package shared

import (
	"fmt"
	"os"
	"sync"

	"github.com/streadway/amqp"
)

type RabbitMQManager struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	mu      sync.Mutex
}

var instance *RabbitMQManager

// GetRabbitMQManager returns a singleton instance of RabbitMQManager
func GetRabbitMQManager() (*RabbitMQManager, error) {
	if instance == nil {
		instance = &RabbitMQManager{}
		err := instance.init()
		if err != nil {
			return nil, err
		}
	}
	return instance, nil
}

// Initialize the RabbitMQ connection and channel
func (r *RabbitMQManager) init() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	var err error
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	if rabbitMQURL == "" {
		rabbitMQURL = "amqp://guest:guest@localhost:5672/" // Default fallback
	}
	r.conn, err = amqp.Dial(rabbitMQURL)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	r.channel, err = r.conn.Channel()
	if err != nil {
		r.conn.Close()
		return fmt.Errorf("failed to open a channel: %w", err)
	}

	return nil
}

// GetChannel returns the RabbitMQ channel
func (r *RabbitMQManager) GetChannel() *amqp.Channel {
	return r.channel
}

// Close cleans up the RabbitMQ connection and channel
func (r *RabbitMQManager) Close() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}
