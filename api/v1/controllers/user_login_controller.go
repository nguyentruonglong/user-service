package controllers

import (
	"log"
	"net/http"
	"time"
	"user-service/api/errors"
	"user-service/api/models"
	"user-service/api/v1/validators"
	"user-service/config"

	firebase "firebase.google.com/go"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// LoginUser logs in a user and returns a Bearer token.
func LoginUser(c *gin.Context, db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) {
	// Parse request body
	var input models.UserLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors.ErrorResponseJSON(c.Writer, errors.ErrInvalidRequestPayload, http.StatusBadRequest)
		return
	}

	// Validate input
	if err := validators.ValidateUserLoginInput(input.Email, input.Password); err != nil {
		errors.ErrorResponseJSON(c.Writer, errors.ErrInvalidInput, http.StatusBadRequest)
		return
	}

	// Perform user authentication
	user, err := authenticateUser(db, input.Email, input.Password)
	if err != nil {
		log.Printf("Authentication failed: %v", err)
		errors.ErrorResponseJSON(c.Writer, errors.ErrAuthenticationFailed, http.StatusUnauthorized)
		return
	}

	// Generate Bearer token with JWT secret key from the configuration
	accessToken, refreshToken, err := generateTokens(user, cfg.JWTSecretKey, cfg.JWTExpiration, cfg.RefreshTokenExpiration, db)
	if err != nil {
		log.Printf("Token generation failed: %v", err)
		errors.ErrorResponseJSON(c.Writer, errors.ErrTokenGenerationFailed, http.StatusInternalServerError)
		return
	}

	// For simplicity, sending a success response with the generated tokens
	successResponse := models.UserLoginResponse{
		Message:      "User logged in successfully",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int(cfg.JWTExpiration.Seconds()), // Convert duration to seconds
	}

	c.JSON(http.StatusOK, successResponse)
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

// generateTokens generates both access and refresh tokens for the given user using JWT and the provided secret key.
func generateTokens(user *models.User, secretKey string, jwtExpiration, refreshTokenExpiration time.Duration, db *gorm.DB) (string, string, error) {
	// Check if a valid refresh token already exists
	existingRefreshToken, err := getValidRefreshToken(user.ID, db)
	if err == nil && existingRefreshToken != "" {
		// If a valid refresh token exists, check if a valid access token exists
		existingAccessToken, err := getValidAccessToken(user.ID, db)
		if err == nil && existingAccessToken != "" {
			// If a valid access token exists, return both tokens
			return existingAccessToken, existingRefreshToken, nil
		}

		// If the access token is not valid, generate a new access token
		return generateAndSaveTokens(user, secretKey, jwtExpiration, refreshTokenExpiration, db)
	}

	// If a valid refresh token does not exist, generate a new refresh token
	refreshToken := generateRefreshToken()

	// Check if a valid access token already exists
	existingAccessToken, err := getValidAccessToken(user.ID, db)
	if err == nil && existingAccessToken != "" {
		// Save the refresh token to the database
		saveRefreshTokenToDatabase(user.ID, refreshToken, time.Now().Add(refreshTokenExpiration), db)
		// If a valid access token exists, return it along with the new refresh token
		return existingAccessToken, refreshToken, nil
	}

	// Save both tokens to the database
	return generateAndSaveTokens(user, secretKey, jwtExpiration, refreshTokenExpiration, db)
}

// generateAndSaveTokens generates new access and refresh tokens and saves them to the database.
func generateAndSaveTokens(user *models.User, secretKey string, jwtExpiration, refreshTokenExpiration time.Duration, db *gorm.DB) (string, string, error) {
	accessToken := generateAccessToken(user, secretKey, jwtExpiration)
	refreshToken := generateRefreshToken()

	// Save the access token to the database
	err := saveAccessTokenToDatabase(user.ID, accessToken, time.Now().Add(jwtExpiration), db)
	if err != nil {
		log.Printf("Failed to save access token to the database: %v", err)
		return "", "", err
	}

	// Save the refresh token to the database
	err = saveRefreshTokenToDatabase(user.ID, refreshToken, time.Now().Add(refreshTokenExpiration), db)
	if err != nil {
		log.Printf("Failed to save refresh token to the database: %v", err)
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// getValidAccessToken retrieves a valid access token for the given user ID from the database.
func getValidAccessToken(userID uint, db *gorm.DB) (string, error) {
	var token models.AccessToken

	// Query for a valid access token for the user
	err := db.Model(&models.AccessToken{}).
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

// generateAccessToken generates a new JWT access token for the given user using JWT and the provided secret key.
func generateAccessToken(user *models.User, secretKey string, jwtExpiration time.Duration) string {
	// Create a new JWT access token
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	// Set claims (payload) for the access token
	claims := jwtToken.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(jwtExpiration).Unix() // Use JWT expiration from the configuration

	// Sign the access token with the provided secret key
	accessToken, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		// Handle the error (e.g., log it)
		log.Printf("Failed to sign access token: %v", err)
		return ""
	}

	return accessToken
}

// saveAccessTokenToDatabase saves the access token to the AccessToken table in the database.
func saveAccessTokenToDatabase(userID uint, tokenString string, expirationTime time.Time, db *gorm.DB) error {
	token := models.AccessToken{
		UserID:         userID, // Set the UserID for the token
		AccessToken:    tokenString,
		ExpirationTime: expirationTime,
	}

	// Save the token to the AccessToken table
	err := db.Create(&token).Error
	if err != nil {
		return err
	}

	return nil
}

// getValidRefreshToken retrieves a valid refresh token for the given user ID from the database.
func getValidRefreshToken(userID uint, db *gorm.DB) (string, error) {
	var token models.RefreshToken

	// Query for a valid refresh token for the user
	err := db.Model(&models.RefreshToken{}).
		Select("refresh_token").
		Where("user_id = ? AND expiration_time > ?", userID, time.Now()).
		Order("expiration_time DESC").
		Limit(1).
		Scan(&token).
		Error
	if err != nil {
		return "", err
	}

	return token.RefreshToken, nil
}

// generateRefreshToken generates a new refresh token for the given user.
func generateRefreshToken() string {
	// Create a new refresh token
	refreshToken := uuid.New().String()

	return refreshToken
}

// saveRefreshTokenToDatabase saves the refresh token to the RefreshToken table in the database.
func saveRefreshTokenToDatabase(userID uint, refreshToken string, expirationTime time.Time, db *gorm.DB) error {
	token := models.RefreshToken{
		UserID:         userID,
		RefreshToken:   refreshToken,
		ExpirationTime: expirationTime,
	}

	// Save the refresh token to the RefreshToken table
	err := db.Create(&token).Error
	if err != nil {
		return err
	}

	return nil
}
