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

	// Parse request body to get the refresh_token
	var input models.UserLogoutInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors.ErrorResponseJSON(c.Writer, errors.ErrInvalidRequestPayload, http.StatusBadRequest)
		return
	}

	// Verify if the refresh_token is provided
	if input.RefreshToken == "" {
		errors.ErrorResponseJSON(c.Writer, errors.ErrInvalidRefreshToken, http.StatusBadRequest)
		return
	}

	// Use a transaction to delete both access and refresh tokens
	err := db.Transaction(func(tx *gorm.DB) error {
		// Delete the access token from the database
		if err := tx.Where("user_id = ?", userIDUint).Delete(&models.AccessToken{}).Error; err != nil {
			return err
		}

		// Delete the associated refresh token as well
		if err := tx.Where("user_id = ? AND refresh_token = ?", userIDUint, input.RefreshToken).Delete(&models.RefreshToken{}).Error; err != nil {
			return err
		}

		return nil
	})

	// Handle transaction errors
	if err != nil {
		errors.ErrorResponseJSON(c.Writer, errors.ErrDatabaseOperationFailed, http.StatusInternalServerError)
		return
	}

	// Sending a success response
	successResponse := models.UserLogoutResponse{
		Message: "User logged out successfully",
	}

	c.JSON(http.StatusOK, successResponse)
}
