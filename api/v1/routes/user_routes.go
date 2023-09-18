// User Routes

package routes

import (
	"net/http"
	"user-service/api/v1/controllers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// RegisterRoutes sets up API routes for the user-related endpoints.
func RegisterRoutes(router *mux.Router, db *gorm.DB) {
	// Example route: Register a user
	router.HandleFunc("/api/v1/register", func(w http.ResponseWriter, r *http.Request) {
		controllers.RegisterUser(w, r, db)
	}).Methods("POST")

	// Add more routes for other user-related actions as needed
}
