// Package controllers

package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"user-service/api/errors"
	"user-service/api/models"
	"user-service/config"

	firebase "firebase.google.com/go"
	"gorm.io/gorm"
)

// LogoutUser logs out a user, effectively invalidating their Bearer token.
func LogoutUser(w http.ResponseWriter, r *http.Request, db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) {
	// Extract the token from the Authorization header
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		errors.ErrorResponseJSON(w, errors.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	// Extract the token from the "Bearer" prefix
	var accessToken string
	tokenParts := strings.Split(authorizationHeader, " ")
	if len(tokenParts) == 2 && tokenParts[0] == "Bearer" {
		accessToken = tokenParts[1]
	} else {
		// No "Bearer" prefix found, consider the whole value as the token
		accessToken = authorizationHeader
	}

	// Check if the token exists and is not expired in the database
	var storedToken models.Token
	err := db.Where("access_token = ? AND expiration_time > ?", accessToken, time.Now()).First(&storedToken).Error
	if err != nil {
		// Token not found or expired
		errors.ErrorResponseJSON(w, errors.ErrUnauthorized, http.StatusUnauthorized)
		return
	}

	// Delete the token from the database
	if err := db.Delete(&storedToken).Error; err != nil {
		errors.ErrorResponseJSON(w, errors.ErrDatabaseOperationFailed, http.StatusInternalServerError)
		return
	}

	// For simplicity, sending a success response
	successResponse := models.UserLogoutResponse{
		Message: "User logged out successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(successResponse)
}
