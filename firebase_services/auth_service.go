// Firebase Authentication Service

package firebase_services

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

// GetFirebaseAuthClient gets the Firebase Auth client.
func GetFirebaseAuthClient(app *firebase.App) (*auth.Client, error) {
	// Retrieve the Firebase Auth client
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}
