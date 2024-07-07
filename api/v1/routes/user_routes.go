// User Routes

package routes

import (
	"user-service/api/middlewares"
	"user-service/api/v1/controllers"
	"user-service/config"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes sets up API routes for the user-related endpoints.
func RegisterRoutes(router *gin.Engine, db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) {
	// Route to register a new user.
	router.POST("/api/v1/register", func(c *gin.Context) {
		controllers.RegisterUser(c, db, firebaseClient, cfg)
	})

	// Route to log in and obtain a Bearer token.
	router.POST("/api/v1/login", func(c *gin.Context) {
		controllers.LoginUser(c, db, firebaseClient, cfg)
	})

	// Route to log out and invalidate Bearer token.
	router.POST("/api/v1/logout", middlewares.AuthMiddleware(db, cfg), func(c *gin.Context) {
		controllers.LogoutUser(c, db, firebaseClient, cfg)
	})

	// Route to send verification SMS.
	router.POST("/api/v1/send-verification-sms", middlewares.AuthMiddleware(db, cfg), func(c *gin.Context) {
		controllers.SendPhoneNumberVerificationCode(c, db, firebaseClient, cfg)
	})

	// Route to send verification email.
	router.POST("/api/v1/send-verification-email", middlewares.AuthMiddleware(db, cfg), func(c *gin.Context) {
		controllers.SendEmailVerificationCode(c, db, firebaseClient, cfg)
	})

	// Route to verify email.
	router.POST("/api/v1/verify-email", middlewares.AuthMiddleware(db, cfg), func(c *gin.Context) {
		controllers.VerifyEmail(c, db, firebaseClient, cfg)
	})
}
