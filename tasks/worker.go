package tasks

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/streadway/amqp"
)

// WorkerConfig contains the configuration for the worker.
type WorkerConfig struct {
	QueueName string
	Processor func(body []byte)
}

// StartWorker starts a generic worker that listens for messages from a specified queue.
func StartWorker(config WorkerConfig) {
	conn, ch, msgs, err := setupRabbitMQ(config.QueueName)
	if err != nil {
		log.Fatalf("Failed to set up RabbitMQ: %v", err)
	}
	defer conn.Close()
	defer ch.Close()

	// Create a channel to receive termination signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Start the worker in a separate goroutine
	go startWorker(msgs, stop, config.Processor)

	// Wait for termination signal
	<-stop
	log.Println("Worker is stopping...")
}

// setupRabbitMQ sets up the RabbitMQ connection, channel, and message queue.
func setupRabbitMQ(queueName string) (*amqp.Connection, *amqp.Channel, <-chan amqp.Delivery, error) {
	// Establish a connection to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, nil, nil, err
	}

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, nil, err
	}

	// Declare a queue
	queue, err := ch.QueueDeclare(
		queueName, // Queue name
		false,     // Durable
		false,     // Delete when unused
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		return nil, nil, nil, err
	}

	// Consume messages from the queue
	msgs, err := ch.Consume(
		queue.Name, // Queue
		"",         // Consumer
		true,       // Auto-acknowledge messages
		false,      // Exclusive
		false,      // No-local
		false,      // No-wait
		nil,        // Args
	)
	if err != nil {
		return nil, nil, nil, err
	}

	return conn, ch, msgs, nil
}

// startWorker processes messages from the queue using the provided processor function.
func startWorker(msgs <-chan amqp.Delivery, stop chan os.Signal, processor func(body []byte)) {
	for {
		select {
		case <-stop:
			return // Stop the worker on termination signal
		case msg := <-msgs:
			// Process the received message using the provided processor function
			processor(msg.Body)
		}
	}
}
