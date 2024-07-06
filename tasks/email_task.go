package tasks

import (
	"encoding/json"
	"fmt"
	"log"
	"user-service/config"
	"user-service/email_services"

	firebase "firebase.google.com/go"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

// EmailTask represents the task data for sending an email.
type EmailTask struct {
	TemplateCode string                 `json:"template_code"`
	Data         map[string]interface{} `json:"data"`
	Recipient    string                 `json:"recipient"`
}

// ProcessEmailTask processes the email tasks from the queue.
func ProcessEmailTask(db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) func(body []byte) {
	return func(body []byte) {
		log.Printf("Received raw message: %s", string(body))

		if len(body) == 0 {
			log.Printf("Received an empty message")
			return
		}

		var task EmailTask
		if err := json.Unmarshal(body, &task); err != nil {
			log.Printf("Error unmarshaling email task: %v, body: %s", err, string(body))
			return
		}

		err := email_services.SendEmail(db, task.TemplateCode, task.Recipient, task.Data)
		if err != nil {
			log.Printf("Error sending email: %v", err)
		}
	}
}

// PublishEmailTask publishes an email task to the RabbitMQ queue.
func PublishEmailTask(queueName string, task EmailTask, cfg *config.AppConfig) error {
	conn, ch, err := setupRabbitMQConnection(cfg)
	if err != nil {
		return err
	}
	defer conn.Close()
	defer ch.Close()

	body, err := json.Marshal(task)
	if err != nil {
		return err
	}

	if len(body) == 0 {
		log.Printf("Attempted to publish an empty task: %v", task)
		return fmt.Errorf("attempted to publish an empty task")
	}

	log.Printf("Publishing email task to queue %s: %s", queueName, string(body))

	err = ch.Publish("", queueName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		return err
	}

	log.Printf("Email task published to queue: %s", queueName)
	return nil
}
