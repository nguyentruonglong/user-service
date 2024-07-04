package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	v1 "user-service/api/v1/routes" // Import v1 package for API routes
	"user-service/config"           // Import config package
	"user-service/database"         // Import database package
	_ "user-service/docs"           // Import docs
	"user-service/firebase_services"

	firebase "firebase.google.com/go" // Import firebase package
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func main() {
	// Parse the command-line argument for the config file path
	configFilePath := flag.String("config", "config/dev_config.yaml", "Path to the configuration file")
	flag.Parse()

	// Load the application configuration
	cfg, err := config.LoadConfig(*configFilePath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the database
	db, err := initDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer database.CloseDB(db)

	// Initialize Firebase app if required
	firebaseClient, err := initFirebase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	// Create a new router using Gin
	router := gin.Default()

	// Register API routes
	v1.RegisterRoutes(router, db, firebaseClient, cfg)

	// Serve Swagger UI in the development environment only
	if isDevConfig(*configFilePath) {
		// Serve Swagger UI
		router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Start the HTTP server
	startServer(router, cfg)
}

// initDatabase initializes the database based on the configuration.
func initDatabase(cfg *config.AppConfig) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	if cfg.GetMultipleDatabasesConfig().GetUseSQLite() {
		db, err = database.InitDB(cfg.GetDatabaseURL())
		if err != nil {
			return nil, err
		}

		// Perform any necessary migrations
		database.AutoMigrateTables(db)
	} else if cfg.GetMultipleDatabasesConfig().GetUsePostgreSQL() {
		// Initialize PostgreSQL database if required
		// db, err = database.InitPostgreSQLDB(cfg.GetDatabaseURL())
		// if err != nil {
		//     return nil, err
		// }
	}

	return db, nil
}

// initFirebase initializes the Firebase app if required by the configuration.
func initFirebase(cfg *config.AppConfig) (*firebase.App, error) {
	if cfg.GetMultipleDatabasesConfig().GetUseRealtimeDatabase() || cfg.GetMultipleDatabasesConfig().GetUseFirestore() {
		ctx := context.Background()
		return firebase_services.InitializeFirebaseApp(ctx, cfg)
	}
	return nil, nil
}

// isDevConfig checks if the configuration file path is for the development environment.
func isDevConfig(configFilePath string) bool {
	return strings.Contains(configFilePath, "dev_config.yaml")
}

// startServer starts the HTTP server using Gin.
func startServer(router *gin.Engine, cfg *config.AppConfig) {
	host := cfg.GetHost()
	port := cfg.GetHTTPPort()
	serverAddr := fmt.Sprintf("%s:%d", host, port)

	log.Printf("Server is running on http://%s\n", serverAddr)
	log.Fatal(router.Run(serverAddr))
}
