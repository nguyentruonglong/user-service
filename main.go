// Main Application Entry Point

package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	v1 "user-service/api/v1/routes" // Import v1 package for API routes
	"user-service/config"           // Import config package
	"user-service/database"         // Import database package
	_ "user-service/docs"           // Import docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
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
	log.Printf("Database URL: %s", cfg.GetDatabaseURL())

	// Initialize the database
	db, err := database.InitDB(cfg.GetDatabaseURL())
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	database.AutoMigrateTables(db)

	defer database.CloseDB(db) // Close the database connection when the server exits

	// Create a new router using Gin
	router := gin.Default()

	// Register API routes
	v1.RegisterRoutes(router, db)

	// Serve Swagger UI in the development environment only
	if isDevConfig {
		// Serve Swagger UI
		router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	fmt.Println(isDevConfig)
	// Set up the server configuration
	host := cfg.GetHost()
	port := cfg.GetHTTPPort()

	// Start the HTTP server
	serverAddr := fmt.Sprintf("%s:%d", host, port)

	// Run the HTTP server using Gin's Run method
	log.Printf("Server is running on http://%s\n", serverAddr)
	log.Fatal(router.Run(serverAddr))
}
