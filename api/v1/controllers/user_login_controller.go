package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"user-service/api/errors"
	"user-service/api/models"
	"user-service/api/v1/validators"
	"user-service/config"

	firebase "firebase.google.com/go"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// LoginUser logs in a user and returns a Bearer token.
func LoginUser(w http.ResponseWriter, r *http.Request, db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) {
	// Parse request body
	var input models.UserLoginInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		errors.ErrorResponseJSON(w, errors.ErrInvalidRequestPayload, http.StatusBadRequest)
		return
	}

	// Validate input
	err = validators.ValidateUserLoginInput(input.Email, input.Password)
	if err != nil {
		errors.ErrorResponseJSON(w, errors.ErrInvalidInput, http.StatusBadRequest)
		return
	}

	// Perform user authentication
	user, err := authenticateUser(db, input.Email, input.Password)
	if err != nil {
		log.Printf("Authentication failed: %v", err)
		errors.ErrorResponseJSON(w, errors.ErrAuthenticationFailed, http.StatusUnauthorized)
		return
	}

	// Generate Bearer token with JWT secret key from the configuration
	token, err := generateToken(user, cfg.GetJWTSecretKey(), cfg.GetJWTExpiration(), db)
	if err != nil {
		log.Printf("Token generation failed: %v", err)
		errors.ErrorResponseJSON(w, errors.ErrTokenGenerationFailed, http.StatusInternalServerError)
		return
	}

	// For simplicity, sending a success response with the generated token
	successResponse := models.UserLoginResponse{
		Message: "User logged in successfully",
		Token:   token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(successResponse)
}

// authenticateUser authenticates a user with the provided email and password.
func authenticateUser(db *gorm.DB, email, password string) (*models.User, error) {
	var user models.User

	// Check if the user exists
	err := db.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	// Compare the provided password with the hashed password in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// generateToken generates a Bearer token for the given user using JWT and the provided secret key.
func generateToken(user *models.User, secretKey string, jwtExpiration time.Duration, db *gorm.DB) (string, error) {
	// Check if a valid token already exists
	existingToken, err := getValidToken(user.ID, db)
	if err == nil && existingToken != "" {
		return existingToken, nil
	}

	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload) for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(jwtExpiration).Unix() // Use JWT expiration from the configuration

	// Sign the token with the provided secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	// Save the token to the database
	err = saveTokenToDatabase(user.ID, tokenString, time.Now().Add(jwtExpiration), db)
	if err != nil {
		// Handle the error (e.g., log it)
		log.Printf("Failed to save token to the database: %v", err)
		return "", err
	}

	return tokenString, nil
}

// getValidToken retrieves a valid token for the given user ID from the database.
func getValidToken(userID uint, db *gorm.DB) (string, error) {
	var token models.Token

	// Query for a valid token for the user
	err := db.Model(&models.Token{}).
		Select("access_token").
		Where("user_id = ? AND expiration_time > ?", userID, time.Now()).
		Order("expiration_time DESC").
		Limit(1).
		Scan(&token).
		Error
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

// saveTokenToDatabase saves the token to the Token table in the database.
func saveTokenToDatabase(userID uint, tokenString string, expirationTime time.Time, db *gorm.DB) error {
	token := models.Token{
		UserID:         userID, // Set the UserID for the token
		AccessToken:    tokenString,
		ExpirationTime: expirationTime,
	}

	// Save the token to the Token table
	err := db.Create(&token).Error
	if err != nil {
		return err
	}

	return nil
}
