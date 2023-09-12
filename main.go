// Main Application Entry Point

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"user-service/config" // Import config package
	"user-service/database"

	"github.com/gorilla/mux"
)

func main() {
	// Parse the command-line argument for config file path
	configFilePath := flag.String("config", "config/dev_config.yml", "Path to the configuration file")
	flag.Parse()

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
	defer database.CloseDB(db) // Close the database connection when the server exits

	// Create a new router using Gorilla Mux.
	router := mux.NewRouter()

	// Define API routes here using router.HandleFunc.

	// Example:
	// router.HandleFunc("/api/v1/register", RegisterHandler).Methods("POST")

	// Set up the server configuration.
	host := cfg.GetHost()
	port := cfg.GetHTTPPort()

	// Start the HTTP server.
	serverAddr := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("Server is running on http://%s\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, router))
}
