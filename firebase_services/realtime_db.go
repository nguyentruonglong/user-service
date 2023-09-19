package firebase_services

import (
	"context"
	"log"

	_ "firebase.google.com/go/db"
)

// SaveDataToRealtimeDB saves data to Firebase Realtime Database.
func SaveDataToRealtimeDB(ctx context.Context, path string, data interface{}) error {
	ref := RealtimeDBClient().NewRef(path)
	if err := ref.Set(ctx, data); err != nil {
		log.Printf("Error saving data to Realtime Database: %v", err)
		return err
	}

	log.Printf("Data saved to Realtime Database at path: %s", path)
	return nil
}
