// User Routes

package routes

import (
	"user-service/api/v1/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes sets up API routes for the user-related endpoints.
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Example route: Register a user
	router.POST("/api/v1/register", func(c *gin.Context) {
		controllers.RegisterUser(c.Writer, c.Request, db)
	})

	// Add more routes for other user-related actions as needed
}
