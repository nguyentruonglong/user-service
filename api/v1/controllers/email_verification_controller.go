package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"user-service/api/errors"
	"user-service/api/models"
	"user-service/config"
	"user-service/tasks"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SendEmailVerificationCode handles the email verification code request
func SendEmailVerificationCode(c *gin.Context, db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) {
	// Get the user ID from the authentication middleware
	userID, _ := c.Get("userID")

	// Fetch user details from the database
	var user models.User
	if err := db.First(&user, userID.(uint)).Error; err != nil {
		errors.ErrorResponseJSON(c.Writer, errors.ErrUserNotFound, http.StatusNotFound)
		return
	}

	// Check if the user email is provided
	if user.Email == "" {
		c.JSON(http.StatusBadRequest, errors.ErrEmailNotProvided)
		return
	}

	// Check if the email is already verified
	if user.IsEmailVerified {
		c.JSON(http.StatusBadRequest, errors.ErrEmailAlreadyVerified)
		return
	}

	// Generate and hash the verification code
	verificationCode := generateEmailVerificationCode()
	hashedCode := hashEmailVerificationCode(verificationCode + user.Email)

	// Update the user's EmailVerificationCode field with the hashed code and set the expiration time
	user.EmailVerificationCode = hashedCode
	user.EmailVerificationExpiry = time.Now().Add(cfg.EmailConfig.VerificationEmailExpiration) // Set the expiration time
	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrDatabaseOperationFailed)
		return
	}

	// Convert the expiration time to a human-readable string
	expiryTimeString := formatDuration(cfg.EmailConfig.VerificationEmailExpiration)

	// Prepare the email task with the verification code
	task := tasks.EmailTask{
		TemplateCode: "EMAIL_VERIFICATION",
		Data: map[string]interface{}{
			"FirstName":        user.FirstName,
			"VerificationCode": verificationCode, // Include the plain code for the email content
			"ExpiryTime":       expiryTimeString,
		},
		Recipient: user.Email,
	}

	if err := tasks.PublishEmailTask("email_queue", task, cfg); err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrEmailTaskPublishingFailed)
		return
	}

	response := models.SendEmailVerificationCodeResponse{
		Message: "Verification email sent",
	}

	c.JSON(http.StatusOK, response)
}

// VerifyEmail handles the email verification using the provided verification code
func VerifyEmail(c *gin.Context, db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) {
	var input models.EmailVerificationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, errors.ErrInvalidEmailVerificationInput)
		return
	}

	// Get the user ID from the authentication middleware
	userID, _ := c.Get("userID")

	// Fetch user details from the database
	var user models.User
	if err := db.First(&user, userID.(uint)).Error; err != nil {
		errors.ErrorResponseJSON(c.Writer, errors.ErrUserNotFound, http.StatusNotFound)
		return
	}

	// Check if the verification code matches
	hashedCode := hashEmailVerificationCode(input.VerificationCode + user.Email)
	if user.EmailVerificationCode != hashedCode {
		c.JSON(http.StatusBadRequest, errors.ErrInvalidVerificationCode)
		return
	}

	// Check if the verification code is expired
	if time.Now().After(user.EmailVerificationExpiry) {
		c.JSON(http.StatusBadRequest, errors.ErrInvalidVerificationCode) // Add appropriate error message for expired code
		return
	}

	// Mark the email as verified
	user.IsEmailVerified = true
	user.EmailVerificationCode = ""            // Clear the verification code
	user.EmailVerificationExpiry = time.Time{} // Clear the expiration time
	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrDatabaseOperationFailed)
		return
	}

	response := models.EmailVerificationResponse{
		Message: "Email verified successfully",
	}

	c.JSON(http.StatusOK, response)
}

// generateEmailVerificationCode generates a random 6-digit verification code.
func generateEmailVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000 // Generate a random 6-digit number
	return fmt.Sprintf("%06d", code)
}

// hashEmailVerificationCode securely hashes the verification code using SHA-256.
func hashEmailVerificationCode(code string) string {
	hash := sha256.New()
	hash.Write([]byte(code))
	hashedCode := hex.EncodeToString(hash.Sum(nil))
	return hashedCode
}

// formatDuration formats a time.Duration into a human-readable string.
func formatDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60

	var parts []string
	if h > 0 {
		parts = append(parts, fmt.Sprintf("%d hours", h))
	}
	if m > 0 {
		parts = append(parts, fmt.Sprintf("%d minutes", m))
	}
	if s > 0 {
		parts = append(parts, fmt.Sprintf("%d seconds", s))
	}

	if len(parts) == 0 {
		return "0 seconds"
	}

	return strings.Join(parts, ", ")
}
