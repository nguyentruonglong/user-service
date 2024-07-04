package firebase_services

import (
	"context"
	"log"
)

// SaveDataToRealtimeDB saves data to Firebase Realtime Database at the specified path.
func SaveDataToRealtimeDB(ctx context.Context, path string, data interface{}) error {
	// Get a reference to the specified path in the Realtime Database
	ref := RealtimeDBClient().NewRef(path)

	// Save the data to the specified path
	if err := ref.Set(ctx, data); err != nil {
		log.Printf("Error saving data to Realtime Database at path %s: %v", path, err)
		return err
	}

	log.Printf("Data successfully saved to Realtime Database at path: %s", path)
	return nil
}
