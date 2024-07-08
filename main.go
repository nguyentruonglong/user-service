package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

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

// @title User Service API
// @version 1.0
// @description This is a sample user service API.
// @host localhost:8080
// @BasePath /api/v1

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
	log.Printf("HTTP Port: %d", cfg.HTTPPort)
	log.Printf("Host: %s", cfg.Host)

	// Initialize SQLite database if enabled
	if cfg.DatabaseConfig.SQLite.Enabled {
		db, err = database.InitSQLiteDB(cfg.DatabaseConfig.SQLite.ConnectionString())
		if err != nil {
			log.Fatalf("Failed to initialize SQLite database: %v", err)
		}
		defer database.CloseDB(db)
	}

	// Initialize PostgreSQL database if enabled
	if cfg.DatabaseConfig.PostgreSQL.Enabled {
		db, err = database.InitPostgreSQLDB(cfg.DatabaseConfig.PostgreSQL.ConnectionString())
		if err != nil {
			log.Fatalf("Failed to initialize PostgreSQL database: %v", err)
		}
		defer database.CloseDB(db)
	}

	// Initialize Firebase if enabled
	if cfg.DatabaseConfig.Firebase.Enabled {
		ctx := context.Background()
		firebaseClient, err = firebase_services.InitializeFirebaseApp(ctx, cfg.DatabaseConfig.Firebase)
		if err != nil {
			log.Fatalf("Failed to initialize Firebase app: %v", err)
		}
	}

	// Seed initial data
	if db != nil {
		if err := database.SeedEmailTemplates(db, firebaseClient, cfg); err != nil {
			log.Fatalf("Failed to seed email templates: %v", err)
		}
	}

	// Create a new router using Gin
	router := gin.Default()

	// Register API routes
	v1.RegisterRoutes(router, db, firebaseClient, cfg)

	// Serve Swagger UI in the development environment only
	if isDevConfig {
		// Swagger endpoint
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Start the HTTP server
	serverAddr := fmt.Sprintf("%s:%d", cfg.Host, cfg.HTTPPort)
	server := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}

	// Start background workers
	ctx, cancel := context.WithCancel(context.Background())
	go tasks.StartAllWorkers(ctx, db, firebaseClient, cfg)

	// Graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		log.Println("Shutting down server...")

		// Create a deadline to wait for.
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()

		// Signal the workers to stop
		cancel()

		// Doesn't block if no connections, but will otherwise wait until the timeout deadline.
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
		}

		log.Println("Server exiting")
	}()

	// Listen and serve
	log.Printf("Server is running on http://%s\n", serverAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %s\n", err)
	}
}
