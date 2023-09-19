// Firebase Authentication Service

package firebase_services

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"golang.org/x/net/context"
)

// GetFirebaseAuthClient gets the Firebase Auth client.
func GetFirebaseAuthClient(app *firebase.App) (*auth.Client, error) {
	return app.Auth(context.Background())
}
