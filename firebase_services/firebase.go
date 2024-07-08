package firebase_services

import (
	"context"
	"log"
	"user-service/config"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

var (
	app              *firebase.App
	realtimeDBClient *db.Client
)

// InitializeFirebaseApp initializes the Firebase app and Realtime Database client.
func InitializeFirebaseApp(ctx context.Context, cfg config.FirebaseConfig) (*firebase.App, error) {
	if app != nil {
		return app, nil
	}

	// Create a Firebase options struct with the provided service account key
	opt := option.WithCredentialsJSON([]byte(cfg.ServiceAccountKey))

	// Create a Firebase Config struct with the necessary fields
	firebaseConfig := &firebase.Config{
		ProjectID:     cfg.ProjectID,
		DatabaseURL:   cfg.DatabaseURL,
		StorageBucket: cfg.StorageBucket,
	}

	// Initialize the Firebase app with the provided options
	var err error
	app, err = firebase.NewApp(ctx, firebaseConfig, opt)
	if err != nil {
		log.Printf("Failed to initialize Firebase app: %v", err)
		return nil, err
	}

	// Initialize the Realtime Database client
	realtimeDBClient, err = app.Database(ctx)
	if err != nil {
		log.Printf("Failed to initialize Realtime Database client: %v", err)
		return nil, err
	}

	return app, nil
}

// RealtimeDBClient returns the initialized Realtime Database client.
func RealtimeDBClient() *db.Client {
	return realtimeDBClient
}
