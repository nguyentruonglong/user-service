package tasks

import (
	"context"
	"log"
	"user-service/config"

	firebase "firebase.google.com/go"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

// WorkerConfig contains the configuration for the worker.
type WorkerConfig struct {
	QueueName string
	Processor func(body []byte)
}

// StartWorker starts a generic worker that listens for messages from a specified queue.
func StartWorker(ctx context.Context, cfg *config.AppConfig, config WorkerConfig) {
	conn, ch, msgs, err := setupRabbitMQ(config.QueueName, cfg)
	if err != nil {
		log.Fatalf("Failed to set up RabbitMQ: %v", err)
	}
	defer conn.Close()
	defer ch.Close()

	// Start the worker in a separate goroutine
	go startWorker(ctx, msgs, config.Processor)

	// Wait for termination signal
	<-ctx.Done()
	log.Println("Worker is stopping...")
}

// setupRabbitMQ sets up the RabbitMQ connection, channel, and message queue.
func setupRabbitMQ(queueName string, cfg *config.AppConfig) (*amqp.Connection, *amqp.Channel, <-chan amqp.Delivery, error) {
	// Establish a connection to RabbitMQ
	conn, err := amqp.Dial(cfg.GetRabbitMQConfig().GetRabbitMQConnectionString())
	if err != nil {
		return nil, nil, nil, err
	}

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, nil, err
	}

	// Declare a queue
	queue, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	// Consume messages from the queue
	msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	return conn, ch, msgs, nil
}

// startWorker processes messages from the queue using the provided processor function.
func startWorker(ctx context.Context, msgs <-chan amqp.Delivery, processor func(body []byte)) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-msgs:
			// Avoid processing empty messages
			if len(msg.Body) == 0 {
				continue
			}
			// Process the received message using the provided processor function
			processor(msg.Body)
		}
	}
}

// StartAllWorkers starts all the defined workers for the application.
func StartAllWorkers(ctx context.Context, db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) {
	go StartWorker(ctx, cfg, WorkerConfig{
		QueueName: "email_queue",
		Processor: ProcessEmailTask(db, firebaseClient, cfg),
	})

	// Add more workers here if needed
}

func setupRabbitMQConnection(cfg *config.AppConfig) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(cfg.GetRabbitMQConfig().GetRabbitMQConnectionString())
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	return conn, ch, nil
}
