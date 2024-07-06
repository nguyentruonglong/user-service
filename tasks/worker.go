package tasks

import (
	"log"
	"os"
	"os/signal"
	"syscall"
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

	// Create a done channel to signal that the worker should stop
	done := make(chan bool)

	// Start the worker in a separate goroutine
	go startWorker(msgs, done, config.Processor)

	// Wait for termination signal
	<-stop
	log.Println("Worker is stopping...")

	// Signal the worker to stop
	close(done)
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
func startWorker(msgs <-chan amqp.Delivery, done chan bool, processor func(body []byte)) {
	for {
		select {
		case <-done:
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
func StartAllWorkers(db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) {
	go StartWorker(WorkerConfig{
		QueueName: "email_queue",
		Processor: ProcessEmailTask(db, firebaseClient, cfg),
	})

	// Add more workers here if needed
}

func setupRabbitMQConnection() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	return conn, ch, nil
}
