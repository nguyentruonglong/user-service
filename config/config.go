package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

// AppConfig holds the application configuration.
type AppConfig struct {
	HTTPPort               int            `mapstructure:"http_port"`
	HTTPSPort              int            `mapstructure:"https_port"`
	Host                   string         `mapstructure:"host"`
	JWTSecretKey           string         `mapstructure:"jwt_secret_key"`
	JWTExpiration          time.Duration  `mapstructure:"jwt_expiration"`
	RefreshTokenExpiration time.Duration  `mapstructure:"refresh_token_expiration"`
	FCMDeviceToken         string         `mapstructure:"fcm_device_token"`
	DatabaseConfig         DatabaseConfig `mapstructure:"database"`
	EmailConfig            EmailConfig    `mapstructure:"email"`
	SMSConfig              SMSConfig      `mapstructure:"sms"`
	RabbitMQConfig         RabbitMQConfig `mapstructure:"rabbitmq"`
}

// DatabaseConfig holds configuration for all enabled databases and settings.
type DatabaseConfig struct {
	SQLite     SQLiteConfig     `mapstructure:"sqlite"`
	PostgreSQL PostgreSQLConfig `mapstructure:"postgresql"`
	Firebase   FirebaseConfig   `mapstructure:"firebase"`
}

// SQLiteConfig holds the SQLite database configuration.
type SQLiteConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	Driver  string `mapstructure:"driver"`
	File    string `mapstructure:"file"`
}

// ConnectionString returns an SQLite connection string built from the config.
func (sqlite *SQLiteConfig) ConnectionString() string {
	if !sqlite.Enabled {
		return "" // Optionally return an empty string if not enabled
	}
	return sqlite.File
}

// PostgreSQLConfig holds the PostgreSQL database configuration.
type PostgreSQLConfig struct {
	Enabled  bool   `mapstructure:"enabled"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

// ConnectionString returns a PostgreSQL connection string built from the config.
func (pg *PostgreSQLConfig) ConnectionString() string {
	if !pg.Enabled {
		return "" // Optionally return an empty string if not enabled
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pg.Host, pg.Port, pg.User, pg.Password, pg.Dbname)
}

// FirebaseConfig holds the Firebase database configuration.
type FirebaseConfig struct {
	Enabled           bool   `mapstructure:"enabled"`
	APIKey            string `mapstructure:"api_key"`
	ProjectID         string `mapstructure:"project_id"`
	StorageBucket     string `mapstructure:"storage_bucket"`
	AuthDomain        string `mapstructure:"auth_domain"`
	DatabaseURL       string `mapstructure:"database_url"`
	ServiceAccountKey string `mapstructure:"service_account_key"`
}

// EmailConfig holds the configuration for email services.
type EmailConfig struct {
	Provider                     string        `mapstructure:"provider"`
	Mailjet                      SMTPConfig    `mapstructure:"mailjet"`
	Sendgrid                     SMTPConfig    `mapstructure:"sendgrid"`
	Generic                      SMTPConfig    `mapstructure:"generic"`
	VerificationEmailExpiration  time.Duration `mapstructure:"verification_email_expiration"`
	PasswordResetEmailExpiration time.Duration `mapstructure:"password_reset_email_expiration"`
}

// SMTPConfig holds the configuration for SMTP services.
type SMTPConfig struct {
	SMTPServer   string `mapstructure:"smtp_server"`
	SMTPPort     int    `mapstructure:"smtp_port"`
	SMTPUser     string `mapstructure:"smtp_user"`
	SMTPPassword string `mapstructure:"smtp_password"`
	SenderEmail  string `mapstructure:"sender_email"`
	SenderName   string `mapstructure:"sender_name"`
}

// SMSConfig holds the SMS service configuration.
type SMSConfig struct {
	TwilioAccountSID  string `mapstructure:"twilio_account_sid"`
	TwilioAuthToken   string `mapstructure:"twilio_auth_token"`
	TwilioPhoneNumber string `mapstructure:"twilio_phone_number"`
}

// RabbitMQConfig holds the RabbitMQ configuration.
type RabbitMQConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

// GetRabbitMQConnectionString returns the RabbitMQ connection string.
func (c *RabbitMQConfig) GetRabbitMQConnectionString() string {
	// Check if the application is running inside a Docker container
	if os.Getenv("RUNNING_IN_DOCKER") == "true" {
		c.Host = "rabbitmq" // Docker service name or container name
	}

	return fmt.Sprintf("amqp://%s:%s@%s:%d/",
		c.Username, c.Password, c.Host, c.Port)
}

// LoadConfig loads the application configuration from a file.
func LoadConfig(configPath string) (*AppConfig, error) {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return &config, nil
}
