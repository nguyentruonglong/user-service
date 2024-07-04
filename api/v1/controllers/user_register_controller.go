package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"user-service/api/errors"
	"user-service/api/models"
	"user-service/api/v1/validators"
	"user-service/config"
	"user-service/utils"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterUser registers a new user and creates an object in the specified database
func RegisterUser(c *gin.Context, db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) {
	// Parse request body
	var input models.UserRegisterInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors.ErrorResponseJSON(c.Writer, errors.ErrInvalidRequestPayload, http.StatusBadRequest)
		return
	}

	// Validate input
	err = validators.ValidateUserRegisterInput(input.Email, input.Password)
	if err != nil {
		errors.ErrorResponseJSON(c.Writer, errors.ErrInvalidInput, http.StatusBadRequest)
		return
	}

	// Start a transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			errors.ErrorResponseJSON(c.Writer, errors.ErrTransactionFailed, http.StatusInternalServerError)
		}
	}()

	// Check if the email already exists
	if emailExistsInDatabase(tx, input.Email) {
		tx.Rollback()
		errors.ErrorResponseJSON(c.Writer, errors.ErrEmailExistsInDatabase, http.StatusConflict)
		return
	}

	// Check if the phone number already exists
	if phoneNumberExistsInDatabase(tx, input.PhoneNumber) {
		tx.Rollback()
		errors.ErrorResponseJSON(c.Writer, errors.ErrPhoneNumberExistsInDatabase, http.StatusConflict)
		return
	}

	// Create a new user with all fields from the API request body
	user := &models.User{
		Email:       input.Email,
		FirstName:   input.FirstName,
		MiddleName:  input.MiddleName,
		LastName:    input.LastName,
		DateOfBirth: input.DateOfBirth,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
		Country:     input.Country,
		Province:    input.Province,
		AvatarURL:   input.AvatarURL,
	}

	// Set the password securely
	err = user.SetPassword(input.Password)
	if err != nil {
		tx.Rollback()
		errors.ErrorResponseJSON(c.Writer, errors.ErrFailedToSetPassword, http.StatusInternalServerError)
		return
	}

	// Save the user to the appropriate database based on the config
	err = saveUserToDatabase(tx, user, cfg)
	if err != nil {
		tx.Rollback()
		errors.ErrorResponseJSON(c.Writer, err, http.StatusInternalServerError)
		return
	}

	// Handle Firebase operations
	err = handleFirebaseOperations(c, user, input, firebaseClient, cfg)
	if err != nil {
		tx.Rollback()
		errors.ErrorResponseJSON(c.Writer, err, http.StatusInternalServerError)
		return
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		errors.ErrorResponseJSON(c.Writer, errors.ErrTransactionFailed, http.StatusInternalServerError)
		return
	}

	// Respond with success message or user data
	successResponse := models.UserRegisterResponse{
		Message: "User registered successfully",
	}
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusCreated)
	c.JSON(http.StatusCreated, successResponse)
}

func saveUserToDatabase(tx *gorm.DB, user *models.User, cfg *config.AppConfig) error {
	if cfg.GetMultipleDatabasesConfig().GetUseSQLite() {
		err := tx.Create(user).Error
		if err != nil {
			return errors.ErrFailedToSaveUserSQLite
		}
	} else if cfg.GetMultipleDatabasesConfig().GetUsePostgreSQL() {
		err := tx.Create(user).Error
		if err != nil {
			return errors.ErrFailedToSaveUserPostgreSQL
		}
	} else {
		return errors.ErrNoValidDatabaseSelected
	}
	return nil
}

func handleFirebaseOperations(c *gin.Context, user *models.User, input models.UserRegisterInput, firebaseClient *firebase.App, cfg *config.AppConfig) error {
	if cfg.GetMultipleDatabasesConfig().GetUseRealtimeDatabase() || cfg.GetMultipleDatabasesConfig().GetUseFirestore() {
		ctx := context.Background()
		client, err := firebaseClient.Database(ctx)
		if err != nil {
			return errors.ErrFailedToGetFirebaseClient
		}

		exists, err := emailExistsInFirebase(client, input.Email)
		if err != nil {
			return errors.ErrFailedToCheckEmailExistence
		}
		if exists {
			return errors.ErrEmailAlreadyExistsOnFirebase
		}

		exists, err = phoneNumberExistsInFirebase(client, input.PhoneNumber)
		if err != nil {
			return errors.ErrFailedToCheckPhoneNumberExistence
		}
		if exists {
			return errors.ErrPhoneNumberAlreadyExistsOnFirebase
		}

		hashedPassword, _ := models.HashPassword(input.Password)

		// Convert the user struct to a map
		userMap := map[string]interface{}{
			"id":                             user.ID,
			"email":                          input.Email,
			"first_name":                     input.FirstName,
			"middle_name":                    utils.GetStringOrDefault(&input.MiddleName, ""),
			"last_name":                      input.LastName,
			"date_of_birth":                  input.DateOfBirth,
			"phone_number":                   input.PhoneNumber,
			"address":                        utils.GetStringOrDefault(&input.Address, ""),
			"country":                        utils.GetStringOrDefault(&input.Country, ""),
			"province":                       utils.GetStringOrDefault(&input.Province, ""),
			"avatar_url":                     utils.GetStringOrDefault(&input.AvatarURL, ""),
			"password_hash":                  hashedPassword,
			"is_active":                      utils.GetBoolOrDefault(nil, true),
			"email_verification_code":        utils.GetStringOrDefault(nil, ""),
			"phone_number_verification_code": utils.GetStringOrDefault(nil, ""),
			"is_email_verified":              utils.GetBoolOrDefault(nil, false),
			"is_phone_number_verified":       utils.GetBoolOrDefault(nil, false),
			"earned_points":                  utils.GetIntOrDefault(nil, 0),
			"extra_info":                     utils.GetOrDefaultJSON(nil, "{}"),
			"created_at":                     time.Now(),
			"updated_at":                     time.Now(),
		}

		if cfg.GetMultipleDatabasesConfig().GetUseRealtimeDatabase() {
			ref := client.NewRef("users")
			newUserRef, err := ref.Push(ctx, userMap)
			if err != nil {
				return errors.ErrFailedToSaveUserFirebaseRTDB
			}

			// Get the unique key generated by Firebase
			firebaseKey := newUserRef.Key
			log.Printf("Firebase Key: %s", firebaseKey)
		}

		if cfg.GetMultipleDatabasesConfig().GetUseFirestore() {
			// TODO: Implement Firestore insertion
		}
	}
	return nil
}

// Function to check if email already exists in the database
func emailExistsInDatabase(db *gorm.DB, email string) bool {
	var count int64
	db.Model(&models.User{}).Where("email = ? AND deleted_at IS NULL", email).Count(&count)
	return count > 0
}

// Function to check if phone number already exists in the database
func phoneNumberExistsInDatabase(db *gorm.DB, phoneNumber string) bool {
	var count int64
	db.Model(&models.User{}).Where("phone_number = ? AND deleted_at IS NULL", phoneNumber).Count(&count)
	return count > 0
}

// Function to check if email already exists in the Firebase Realtime Database
func emailExistsInFirebase(client *db.Client, email string) (bool, error) {
	ctx := context.Background()

	// Get a reference to the users node in the Firebase Realtime Database
	ref := client.NewRef("users")

	// Query to check if the email exists
	query := ref.OrderByChild("email").EqualTo(email)

	// Get the result of the query
	results, err := query.GetOrdered(ctx)
	if err != nil {
		// Print or log the error
		fmt.Println("Error querying Firebase:", err)
		return false, err
	}

	// Check if there are any results
	if results == nil {
		// Handle the case of nil results
		fmt.Println("Results are nil.")
		return false, nil
	}

	// Log the results for debugging
	fmt.Println("Query Results:", results)

	return len(results) > 0, nil
}

// Function to check if phone number already exists in the Firebase Realtime Database
func phoneNumberExistsInFirebase(client *db.Client, phoneNumber string) (bool, error) {
	ctx := context.Background()

	// Get a reference to the users node in the Firebase Realtime Database
	ref := client.NewRef("users")

	// Query to check if the phone number exists
	query := ref.OrderByChild("phone_number").EqualTo(phoneNumber)

	// Get the result of the query
	results, err := query.GetOrdered(ctx)
	if err != nil {
		// Print or log the error
		fmt.Println("Error querying Firebase:", err)
		return false, err
	}

	// Check if there are any results
	if results == nil {
		// Handle the case of nil results
		fmt.Println("Results are nil.")
		return false, nil
	}

	// Log the results for debugging
	fmt.Println("Query Results:", results)

	return len(results) > 0, nil
}
