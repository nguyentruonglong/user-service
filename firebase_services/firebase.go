package firebase_services

import (
	"context"
	"log"
	"user-service/config"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

var app *firebase.App
var realtimeDBClient *db.Client

func InitializeFirebaseApp(ctx context.Context, cfg *config.AppConfig) (*firebase.App, error) {
	var err error

	if app == nil {
		// Create a Firebase options struct with the provided service account key
		opt := option.WithCredentialsJSON([]byte(cfg.FirebaseConfig.ServiceAccountKey))

		// Create a Firebase Config struct with the necessary fields
		firebaseConfig := &firebase.Config{
			ProjectID:     cfg.FirebaseConfig.ProjectID,
			DatabaseURL:   cfg.FirebaseConfig.DatabaseURL,
			StorageBucket: cfg.FirebaseConfig.StorageBucket,
		}

		// Initialize the Firebase app with the provided options
		app, err = firebase.NewApp(ctx, firebaseConfig, opt)
		if err != nil {
			log.Fatalf("Failed to initialize Firebase app: %v", err)
			return nil, err
		}

		// Initialize the Realtime Database client
		realtimeDBClient, err = app.Database(ctx)
		if err != nil {
			log.Fatalf("Failed to initialize Realtime Database client: %v", err)
			return nil, err
		}
	}

	return app, nil
}

func RealtimeDBClient() *db.Client {
	return realtimeDBClient
}
