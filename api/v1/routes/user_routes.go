// User Routes

package routes

import (
	"user-service/api/v1/controllers"

	"user-service/config"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes sets up API routes for the user-related endpoints.
func RegisterRoutes(router *gin.Engine, db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) {
	// Example route: Register a user
	router.POST("/api/v1/register", func(c *gin.Context) {
		controllers.RegisterUser(c.Writer, c.Request, db, firebaseClient, cfg)
	})

	// Add more routes for other user-related actions as needed
}
