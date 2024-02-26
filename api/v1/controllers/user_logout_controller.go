// Package controllers

package controllers

import (
	"net/http"
	"user-service/api/errors"
	"user-service/api/models"
	"user-service/config"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LogoutUser logs out a user, effectively invalidating their Bearer token.
func LogoutUser(c *gin.Context, db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) {
	// Retrieve user ID from the context
	userID, exists := c.Get("userID")
	if !exists {
		errors.ErrorResponseJSON(c.Writer, errors.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	// Convert userID to uint
	userIDUint, ok := userID.(uint)
	if !ok {
		errors.ErrorResponseJSON(c.Writer, errors.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	// Delete the token from the database
	if err := db.Where("user_id = ?", userIDUint).Delete(&models.Token{}).Error; err != nil {
		errors.ErrorResponseJSON(c.Writer, errors.ErrDatabaseOperationFailed, http.StatusInternalServerError)
		return
	}

	// For simplicity, sending a success response
	successResponse := models.UserLogoutResponse{
		Message: "User logged out successfully",
	}

	c.JSON(http.StatusOK, successResponse)
}
