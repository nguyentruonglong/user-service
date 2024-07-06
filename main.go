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
	"user-service/tasks"

	firebase "firebase.google.com/go" // Import firebase package
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB
	var firebaseClient *firebase.App

	// Parse the command-line argument for the config file path
	configFilePath := flag.String("config", "config/dev_config.yaml", "Path to the configuration file")
	flag.Parse()

	// Check if the config file path contains "dev_config.yaml"
	isDevConfig := strings.Contains(*configFilePath, "dev_config.yaml")

	// Load the application configuration
	cfg, err := config.LoadConfig(*configFilePath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Access configuration values
	log.Printf("HTTP Port: %d", cfg.GetHTTPPort())
	log.Printf("Host: %s", cfg.GetHost())

	if cfg.GetMultipleDatabasesConfig().GetUseSQLite() {
		// Explicitly open and close the database to ensure it's created
		db, err = database.InitDB(cfg)
		if err != nil {
			log.Fatalf("Failed to initialize the database: %v", err)
		}
		defer database.CloseDB(db)
	} else if cfg.GetMultipleDatabasesConfig().GetUsePostgreSQL() {
	}

	if cfg.GetMultipleDatabasesConfig().GetUseRealtimeDatabase() {
		// Initialize Firebase app
		ctx := context.Background()
		firebaseClient, err = firebase_services.InitializeFirebaseApp(ctx, cfg)
		if err != nil {
			log.Fatalf("Failed to initialize Firebase app: %v", err)
		}
	}

	// Seed initial data
	if err := database.SeedEmailTemplates(db, firebaseClient, cfg); err != nil {
		log.Fatalf("Failed to seed email templates: %v", err)
	}

	// Start all workers
	tasks.StartAllWorkers(db, firebaseClient, cfg)

	// Create a new router using Gin
	router := gin.Default()

	// Register API routes
	v1.RegisterRoutes(router, db, firebaseClient, cfg)

	// Serve Swagger UI in the development environment only
	if isDevConfig {
		// Serve Swagger UI
		router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Set up the server configuration
	host := cfg.GetHost()
	port := cfg.GetHTTPPort()

	// Start the HTTP server
	serverAddr := fmt.Sprintf("%s:%d", host, port)

	// Run the HTTP server using Gin's Run method
	log.Printf("Server is running on http://%s\n", serverAddr)
	log.Fatal(router.Run(serverAddr))
}
