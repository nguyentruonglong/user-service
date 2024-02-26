// Configuration Module

package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// AppConfig holds the application configuration.
type AppConfig struct {
	HTTPPort                int                     `mapstructure:"http_port"`
	HTTPSPort               int                     `mapstructure:"https_port"`
	Host                    string                  `mapstructure:"host"`
	JWTSecretKey            string                  `mapstructure:"jwt_secret_key"`
	JWTExpiration           time.Duration           `mapstructure:"jwt_expiration"`
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

// General Configuration Functions

// GetHTTPPort returns the configured HTTP port.
func (c *AppConfig) GetHTTPPort() int {
	return c.HTTPPort
}

// GetHost returns the configured host.
func (c *AppConfig) GetHost() string {
	if c.Host == "" {
		return "localhost"
	}
	return c.Host
}

// GetJWTSecretKey returns the configured JWT secret key.
func (c *AppConfig) GetJWTSecretKey() string {
	return c.JWTSecretKey
}

// GetJWTExpiration returns the configured JWT expiration duration.
func (c *AppConfig) GetJWTExpiration() time.Duration {
	return c.JWTExpiration
}

// Database Configuration Functions

// GetDatabaseURL returns the database URL based on the configured database type.
func (c *AppConfig) GetDatabaseURL() string {
	if c.GetMultipleDatabasesConfig().GetUseSQLite() {
		return c.GetSQLiteConfig().GetConnectionString()
	} else if c.GetMultipleDatabasesConfig().GetUsePostgreSQL() {
		return c.GetPostgreSQLConfig().GetPostgreSQLConnectionString()
	}

	// Default to SQLite if neither is specified
	return c.SQLiteConfig.ConnectionString
}

// GetMultipleDatabasesConfig returns a pointer to the configuration for multiple databases.
func (c *AppConfig) GetMultipleDatabasesConfig() *MultipleDatabasesConfig {
	return &c.MultipleDatabasesConfig
}

// GetUseSQLite returns the UseSQLite configuration.
func (c *MultipleDatabasesConfig) GetUseSQLite() bool {
	return c.UseSQLite
}

// GetUsePostgreSQL returns the UsePostgreSQL configuration.
func (c *MultipleDatabasesConfig) GetUsePostgreSQL() bool {
	return c.UsePostgreSQL
}

// GetUseRealtimeDatabase returns the UseRealtimeDatabase configuration.
func (c *MultipleDatabasesConfig) GetUseRealtimeDatabase() bool {
	return c.UseRealtimeDatabase
}

// GetUseFirestore returns the UseFirestore configuration.
func (c *MultipleDatabasesConfig) GetUseFirestore() bool {
	return c.UseFirestore
}

// SQLite Configuration Functions

// GetSQLiteConfig returns the SQLite configuration.
func (c *AppConfig) GetSQLiteConfig() *SQLiteConfig {
	return &c.SQLiteConfig
}

// GetSQLiteConfig returns the SQLite Connection String configuration.
func (c *SQLiteConfig) GetConnectionString() string {
	return c.ConnectionString
}

// PostgreSQL Configuration Functions

// GetPostgreSQLConfig returns the PostgreSQL configuration.
func (c *AppConfig) GetPostgreSQLConfig() *PostgreSQLConfig {
	return &c.PostgreSQLConfig
}

// GetHost returns the PostgreSQL host configuration.
func (c *PostgreSQLConfig) GetHost() string {
	return c.Host
}

// GetPort returns the PostgreSQL port configuration.
func (c *PostgreSQLConfig) GetPort() int {
	return c.Port
}

// GetUser returns the PostgreSQL user configuration.
func (c *PostgreSQLConfig) GetUser() string {
	return c.User
}

// GetPassword returns the PostgreSQL password configuration.
func (c *PostgreSQLConfig) GetPassword() string {
	return c.Password
}

// GetDbname returns the PostgreSQL dbname configuration.
func (c *PostgreSQLConfig) GetDbname() string {
	return c.Dbname
}

// GetPostgreSQLConnectionString returns the PostgreSQL connection string configuration.
func (c *PostgreSQLConfig) GetPostgreSQLConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.GetHost(), c.GetPort(), c.GetUser(), c.GetPassword(), c.GetDbname())
}

// Firebase Configuration Functions

// GetFirebaseConfig returns the Firebase configuration.
func (c *AppConfig) GetFirebaseConfig() *FirebaseConfig {
	return &c.FirebaseConfig
}

// GetFirebaseURL returns the Firebase database URL.
func (c *AppConfig) GetFirebaseURL() string {
	return c.FirebaseConfig.DatabaseURL
}

// GetAPIKey returns the Firebase API Key configuration.
func (c *FirebaseConfig) GetAPIKey() string {
	return c.APIKey
}

// GetProjectID returns the Firebase Project ID configuration.
func (c *FirebaseConfig) GetProjectID() string {
	return c.ProjectID
}

// GetStorageBucket returns the Firebase Storage Bucket configuration.
func (c *FirebaseConfig) GetStorageBucket() string {
	return c.StorageBucket
}

// GetAuthDomain returns the Firebase Auth Domain configuration.
func (c *FirebaseConfig) GetAuthDomain() string {
	return c.AuthDomain
}

// GetDatabaseURL returns the Firebase Database URL configuration.
func (c *FirebaseConfig) GetDatabaseURL() string {
	return c.DatabaseURL
}

// GetServiceAccountKey returns the Firebase Service Account Key configuration.
func (c *FirebaseConfig) GetServiceAccountKey() string {
	return c.ServiceAccountKey
}

// Email Configuration Functions

// GetEmailConfig returns the email service configuration.
func (c *AppConfig) GetEmailConfig() *EmailConfig {
	return &c.EmailConfig
}

// GetSMTPServer returns the SMTP server configuration.
func (c *EmailConfig) GetSMTPServer() string {
	return c.SMTPServer
}

// GetSMTPPort returns the SMTP port configuration.
func (c *EmailConfig) GetSMTPPort() int {
	return c.SMTPPort
}

// GetSMTPUser returns the SMTP user configuration.
func (c *EmailConfig) GetSMTPUser() string {
	return c.SMTPUser
}

// GetSMTPPassword returns the SMTP password configuration.
func (c *EmailConfig) GetSMTPPassword() string {
	return c.SMTPPassword
}

// SMS Configuration Functions

// GetSMSConfig returns the SMS service configuration.
func (c *AppConfig) GetSMSConfig() *SMSConfig {
	return &c.SMSConfig
}

// GetTwilioAccountSID returns the Twilio Account SID configuration.
func (c *SMSConfig) GetTwilioAccountSID() string {
	return c.TwilioAccountSID
}

// GetTwilioAuthToken returns the Twilio Auth Token configuration.
func (c *SMSConfig) GetTwilioAuthToken() string {
	return c.TwilioAuthToken
}

// GetTwilioPhoneNumber returns the Twilio Phone Number configuration.
func (c *SMSConfig) GetTwilioPhoneNumber() string {
	return c.TwilioPhoneNumber
}
