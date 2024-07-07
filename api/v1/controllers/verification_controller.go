package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
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
	hashedCode := hashEmailVerificationCode(verificationCode)

	// Update the user's EmailVerificationCode field with the hashed code
	user.EmailVerificationCode = hashedCode
	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrDatabaseOperationFailed)
		return
	}

	// Prepare the email task with the verification code
	task := tasks.EmailTask{
		TemplateCode: "EMAIL_VERIFICATION",
		Data: map[string]interface{}{
			"FirstName":        user.FirstName,
			"VerificationCode": verificationCode, // Include the plain code for the email content
		},
		Recipient: user.Email,
	}

	if err := tasks.PublishEmailTask("email_queue", task, cfg); err != nil {
		c.JSON(http.StatusInternalServerError, errors.ErrEmailTaskPublishingFailed)
		return
	}

	response := models.EmailVerificationResponse{
		Message: "Verification email sent",
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
