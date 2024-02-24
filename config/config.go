// Configuration Module

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// AppConfig holds the application configuration.
type AppConfig struct {
	HTTPPort                int                     `mapstructure:"http_port"`
	HTTPSPort               int                     `mapstructure:"https_port"`
	Host                    string                  `mapstructure:"host"`
	JWTSecretKey            string                  `mapstructure:"jwt_secret_key"`
	MultipleDatabasesConfig MultipleDatabasesConfig `mapstructure:"multiple_databases"`
	SQLiteConfig            SQLiteConfig            `mapstructure:"sqlite"`
	PostgreSQLConfig        PostgreSQLConfig        `mapstructure:"postgresql"`
	FirebaseConfig          FirebaseConfig          `mapstructure:"firebase"`
	EmailConfig             EmailConfig             `mapstructure:"email"`
	SMSConfig               SMSConfig               `mapstructure:"sms"`
}

// MultipleDatabasesConfig holds the multiple databases configuration.
type MultipleDatabasesConfig struct {
	UseRealtimeDatabase bool `mapstructure:"use_realtime_database"`
	UseFirestore        bool `mapstructure:"use_firestore"`
	UseSQLite           bool `mapstructure:"use_sqlite"`
	UsePostgreSQL       bool `mapstructure:"use_postgresql"`
}

// SQLiteConfig holds the SQLite configuration.
type SQLiteConfig struct {
	Driver           string `mapstructure:"driver"`
	ConnectionString string `mapstructure:"connection_string"`
}

// PostgreSQLConfig holds the PostgreSQL configuration.
type PostgreSQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

// FirebaseConfig holds the Firebase configuration.
type FirebaseConfig struct {
	APIKey            string `mapstructure:"api_key"`
	ProjectID         string `mapstructure:"project_id"`
	StorageBucket     string `mapstructure:"storage_bucket"`
	AuthDomain        string `mapstructure:"auth_domain"`
	DatabaseURL       string `mapstructure:"database_url"`
	ServiceAccountKey string `mapstructure:"service_account_key"`
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
	if c.MultipleDatabasesConfig.UseSQLite {
		return c.SQLiteConfig.ConnectionString
	} else if c.MultipleDatabasesConfig.UsePostgreSQL {
		return c.PostgreSQLConnectionString()
	}
	// Default to SQLite if neither is specified
	return c.SQLiteConfig.ConnectionString
}

func (c *AppConfig) GetFirebaseURL() string {
	return c.FirebaseConfig.DatabaseURL
}

func (c *AppConfig) GetEmailServiceConfig() EmailConfig {
	return c.EmailConfig
}

func (c *AppConfig) GetSMSServiceConfig() SMSConfig {
	return c.SMSConfig
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

func (c *AppConfig) GetJWTSecretKey() string {
	return c.JWTSecretKey
}

func (c *AppConfig) PostgreSQLConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.PostgreSQLConfig.Host, c.PostgreSQLConfig.Port, c.PostgreSQLConfig.User, c.PostgreSQLConfig.Password, c.PostgreSQLConfig.Dbname)
}
