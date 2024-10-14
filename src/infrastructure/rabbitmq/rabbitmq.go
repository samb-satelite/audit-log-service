package rabbitmq

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

var RabbitMQURL string
var RabbitMQInstance *RabbitMQ

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// InitRabbitMQ initializes RabbitMQ connection
func InitRabbitMQ() error {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: Could not load .env file, using environment variables from the host")
	}

	// Get RABBITMQ_URL from environment variable
	RabbitMQURL = os.Getenv("RABBITMQ_URL")
	if RabbitMQURL == "" {
		log.Println("Warning: RABBITMQ_URL must be set in the environment variables or .env file")
		return nil // Don't exit if the URL is not set
	}

	var err error
	RabbitMQInstance, err = NewRabbitMQ(RabbitMQURL)
	if err != nil {
		log.Println("Warning: Failed to initialize RabbitMQ, retrying...")
		if retryConnection() {
			log.Println("Successfully connected to RabbitMQ after retry")
		} else {
			log.Println("Warning: Could not connect to RabbitMQ after multiple attempts")
		}
		return nil // Allow the app to continue running
	}

	RabbitMQInstance.Channel, err = RabbitMQInstance.Connection.Channel()
	if err != nil {
		log.Println("Warning: Failed to open a channel:", err)
		return nil // Allow the app to continue running
	}

	log.Println("Connected to RabbitMQ")
	return nil
}

// NewRabbitMQ creates a RabbitMQ instance
func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return &RabbitMQ{Connection: conn}, nil
}

// retryConnection attempts to connect to RabbitMQ with retries
func retryConnection() bool {
	for attempts := 0; attempts < 5; attempts++ {
		time.Sleep(2 * time.Second) // Wait before retrying
		var err error
		RabbitMQInstance, err = NewRabbitMQ(RabbitMQURL)
		if err == nil {
			return true // Successfully connected
		}
		log.Printf("Retrying RabbitMQ connection: attempt %d failed: %s\n", attempts+1, err)
	}
	return false // Failed after multiple attempts
}

// Close closes the RabbitMQ connection
func (r *RabbitMQ) Close() {
	if err := r.Channel.Close(); err != nil {
		log.Fatalf("Failed to close RabbitMQ channel: %s", err)
	}
	if err := r.Connection.Close(); err != nil {
		log.Fatalf("Failed to close RabbitMQ connection: %s", err)
	}
}

func (r *RabbitMQ) DeclareQueue(queue string) error {
	_, err := r.Channel.QueueDeclare(
		queue, // Name of the queue
		true,  // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Additional arguments
	)
	return err
}
