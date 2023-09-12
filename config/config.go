// Configuration Module

package config

import (
	"github.com/spf13/viper"
)

// AppConfig holds the application configuration.
type AppConfig struct {
	HTTPPort     int    `mapstructure:"http_port"`
	HTTPSPort    int    `mapstructure:"https_port"`
	Host         string `mapstructure:"host"`
	Database     DatabaseConfig
	Firebase     FirebaseConfig
	EmailService EmailConfig `mapstructure:"email"`
	SMSService   SMSConfig   `mapstructure:"sms"`
}

// DatabaseConfig holds the database configuration.
type DatabaseConfig struct {
	Driver           string `mapstructure:"driver"`
	ConnectionString string `mapstructure:"connection_string"`
}

// FirebaseConfig holds the Firebase configuration.
type FirebaseConfig struct {
	APIKey        string `mapstructure:"api_key"`
	ProjectID     string `mapstructure:"project_id"`
	StorageBucket string `mapstructure:"storage_bucket"`
	AuthDomain    string `mapstructure:"auth_domain"`
	DatabaseURL   string `mapstructure:"database_url"`
}

// EmailConfig holds the email service configuration.
type EmailConfig struct {
	SMTPServer   string `mapstructure:"smtp_server"`
	SMTPPort     int    `mapstructure:"smtp_port"`
	SMTPUser     string `mapstructure:"smtp_user"`
	SMTPPassword string `mapstructure:"smtp_password"`
}

// SMSConfig holds the SMS service configuration.
type SMSConfig struct {
	TwilioAccountSID  string `mapstructure:"twilio_account_sid"`
	TwilioAuthToken   string `mapstructure:"twilio_auth_token"`
	TwilioPhoneNumber string `mapstructure:"twilio_phone_number"`
}

// LoadConfig loads the application configuration.
func LoadConfig(configFilePath string) (*AppConfig, error) {
	// Initialize a new AppConfig
	cfg := &AppConfig{}

	// Use Viper to read the config file
	viper.SetConfigFile(configFilePath)

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// Unmarshal the configuration into the AppConfig struct
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *AppConfig) GetDatabaseURL() string {
	return c.Database.ConnectionString
}

func (c *AppConfig) GetFirebaseURL() string {
	return c.Firebase.DatabaseURL
}

func (c *AppConfig) GetEmailServiceConfig() EmailConfig {
	return c.EmailService
}

func (c *AppConfig) GetSMSServiceConfig() SMSConfig {
	return c.SMSService
}

func (c *AppConfig) GetHTTPPort() int {
	return c.HTTPPort
}

func (c *AppConfig) GetHost() string {
	if c.Host == "" {
		return "localhost"
	}
	return c.Host
}
