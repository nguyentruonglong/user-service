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
	// Start a transaction
	tx := db.Begin()

	// Rollback the transaction if any error occurs
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			errors.ErrorResponseJSON(c.Writer, errors.ErrTransactionFailed, http.StatusInternalServerError)
		}
	}()

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

	// Check if the email already exists
	if emailExistsInDatabase(tx, input.Email) {
		errors.ErrorResponseJSON(c.Writer, errors.ErrEmailExistsInDatabase, http.StatusConflict)
		return
	}

	// Check if the phone number already exists
	if phoneNumberExistsInDatabase(tx, input.PhoneNumber) {
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
	if cfg.GetMultipleDatabasesConfig().GetUseSQLite() {
		err = tx.Create(user).Error
		if err != nil {
			tx.Rollback()
			errors.ErrorResponseJSON(c.Writer, errors.ErrFailedToSaveUserSQLite, http.StatusInternalServerError)
			return
		}
	} else if cfg.GetMultipleDatabasesConfig().GetUsePostgreSQL() {
		err = tx.Create(user).Error
		if err != nil {
			tx.Rollback()
			errors.ErrorResponseJSON(c.Writer, errors.ErrFailedToSaveUserPostgreSQL, http.StatusInternalServerError)
			return
		}
		// TODO: Insert into PostgreSQL database
	} else {
		errors.ErrorResponseJSON(c.Writer, errors.ErrNoValidDatabaseSelected, http.StatusBadRequest)
		return
	}

	// Initialize Firebase app if configured to use Firebase
	if cfg.GetMultipleDatabasesConfig().GetUseRealtimeDatabase() || cfg.GetMultipleDatabasesConfig().GetUseFirestore() {
		ctx := context.Background()
		// Get a Firebase database client
		client, err := firebaseClient.Database(ctx)
		if err != nil {
			tx.Rollback()
			errors.ErrorResponseJSON(c.Writer, errors.ErrFailedToGetFirebaseClient, http.StatusInternalServerError)
			return
		}

		exists, err := emailExistsInFirebase(client, input.Email)
		if err != nil {
			tx.Rollback()
			// Handle case where there was an error checking email existence in Firebase
			errors.ErrorResponseJSON(c.Writer, errors.ErrFailedToCheckEmailExistence, http.StatusInternalServerError)
			return
		}

		if exists {
			tx.Rollback()
			// Handle case where email already exists in Firebase
			errors.ErrorResponseJSON(c.Writer, errors.ErrEmailAlreadyExistsOnFirebase, http.StatusConflict)
			return
		}

		// Check if the phone number already exists in Firebase
		exists, err = phoneNumberExistsInFirebase(client, input.PhoneNumber)
		if err != nil {
			tx.Rollback()
			// Handle case where there was an error checking phone number existence in Firebase
			errors.ErrorResponseJSON(c.Writer, errors.ErrFailedToCheckPhoneNumberExistence, http.StatusInternalServerError)
			return
		}

		if exists {
			tx.Rollback()
			// Handle case where phone number already exists in Firebase
			errors.ErrorResponseJSON(c.Writer, errors.ErrPhoneNumberAlreadyExistsOnFirebase, http.StatusConflict)
			return
		}

		hashedPassword, _ := models.HashPassword(input.Password) // Hash the input password

		// Convert the user struct to a map
		userMap := map[string]interface{}{
			"id":                             user.ID,
			"email":                          input.Email,
			"first_name":                     input.FirstName,
			"middle_name":                    utils.GetStringOrDefault(&input.MiddleName, ""), // Add default middle_name value
			"last_name":                      input.LastName,
			"date_of_birth":                  input.DateOfBirth,
			"phone_number":                   input.PhoneNumber,
			"address":                        utils.GetStringOrDefault(&input.Address, ""),   // Add default address value
			"country":                        utils.GetStringOrDefault(&input.Country, ""),   // Add default country value
			"province":                       utils.GetStringOrDefault(&input.Province, ""),  // Add default province value
			"avatar_url":                     utils.GetStringOrDefault(&input.AvatarURL, ""), // Add default avatar_url value
			"password_hash":                  hashedPassword,
			"is_active":                      utils.GetBoolOrDefault(nil, true),  // Add default is_active value
			"email_verification_code":        utils.GetStringOrDefault(nil, ""),  // Add default email_verification_code value
			"phone_number_verification_code": utils.GetStringOrDefault(nil, ""),  // Add default phone_number_verification_code value
			"is_email_verified":              utils.GetBoolOrDefault(nil, false), // Add default is_email_verified value
			"is_phone_number_verified":       utils.GetBoolOrDefault(nil, false), // Add default is_phone_number_verified value
			"earned_points":                  utils.GetIntOrDefault(nil, 0),      // Add default earned_points value
			"extra_info":                     utils.GetOrDefaultJSON(nil, "{}"),  // Add default extra_info value
			"created_at":                     time.Now(),                         // Add default created_at value
			"updated_at":                     time.Now(),                         // Add default updated_at value
		}

		if cfg.GetMultipleDatabasesConfig().GetUseRealtimeDatabase() {
			// Push the user data to Firebase Realtime Database
			ref := client.NewRef("users")
			newUserRef, err := ref.Push(context.Background(), userMap)
			if err != nil {
				tx.Rollback()
				errors.ErrorResponseJSON(c.Writer, errors.ErrFailedToSaveUserFirebaseRTDB, http.StatusInternalServerError)
				return
			}

			// Get the unique key generated by Firebase
			firebaseKey := newUserRef.Key
			log.Printf("Firebase Key: %s", firebaseKey)
		}

		if cfg.GetMultipleDatabasesConfig().GetUseFirestore() {
			// TODO: Implement Firestore insertion
		}
	}

	// Commit the transaction
	tx.Commit()

	// Respond with success message or user data
	successResponse := models.UserRegisterResponse{
		Message: "User registered successfully",
	}
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusCreated)
	c.JSON(http.StatusCreated, successResponse)
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
