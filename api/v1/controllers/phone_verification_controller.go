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

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"gorm.io/gorm"
)

// SendPhoneNumberVerificationCode sends a verification SMS to the user's phone number.
func SendPhoneNumberVerificationCode(c *gin.Context, db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) {
	// Get the user ID from the authentication middleware
	userID, _ := c.Get("userID")

	// Fetch user details from the database
	var user models.User
	if err := db.First(&user, userID.(uint)).Error; err != nil {
		errors.ErrorResponseJSON(c.Writer, errors.ErrUserNotFound, http.StatusNotFound)
		return
	}

	// Check if the phone number does not exist in the database.
	if user.PhoneNumber == "" {
		errors.ErrorResponseJSON(c.Writer, errors.ErrPhoneNumberNotFoundInDatabase, http.StatusBadRequest)
		return
	}

	// Generate a random verification code
	phoneNumberVerificationCode := generatePhoneNumberVerificationCode()

	// Send the verification SMS using Firebase Cloud Messaging (FCM)
	if err := sendVerificationSMS(user.PhoneNumber, phoneNumberVerificationCode, cfg); err != nil {
		errors.ErrorResponseJSON(c.Writer, errors.ErrSMSFailure, http.StatusInternalServerError)
		return
	}

	// Hash the verification code securely before saving to the database
	hashedCode := hashVerificationCode(phoneNumberVerificationCode + user.PhoneNumber)

	// Save the verification code to the user record in the database
	user.PhoneNumberVerificationCode = hashedCode
	if err := db.Save(&user).Error; err != nil {
		errors.ErrorResponseJSON(c.Writer, errors.ErrDatabaseOperationFailed, http.StatusInternalServerError)
		return
	}

	// Respond with success
	successResponse := models.PhoneNumberVerificationResponse{
		Message: "Verification SMS sent successfully",
	}
	c.JSON(http.StatusOK, successResponse)
}

// generatePhoneNumberVerificationCode generates a random 6-digit verification code.
func generatePhoneNumberVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000 // Generate a random 6-digit number
	return fmt.Sprintf("%06d", code)
}

// hashVerificationCode securely hashes the verification code using SHA-256.
func hashVerificationCode(code string) string {
	hash := sha256.New()
	hash.Write([]byte(code))
	hashedCode := hex.EncodeToString(hash.Sum(nil))
	return hashedCode
}

// sendVerificationSMS sends a verification SMS-like message using Twilio.
func sendVerificationSMS(phoneNumber, verificationCode string, cfg *config.AppConfig) error {
	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo(phoneNumber)
	params.SetFrom(cfg.GetSMSConfig().GetTwilioPhoneNumber())
	params.SetBody(verificationCode)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		return err
	}

	return nil
}
