package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"user-service/api/errors"
	"user-service/api/models"
	"user-service/api/v1/validators"
	"user-service/config"
	"user-service/utils"

	firebase "firebase.google.com/go"

	"gorm.io/gorm"
)

// RegisterUser registers a new user and creates an object in the specified database
func RegisterUser(w http.ResponseWriter, r *http.Request, db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) {
	// Start a transaction
	tx := db.Begin()

	// Rollback the transaction if any error occurs
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			errors.ErrorResponseJSON(w, errors.ErrTransactionFailed, http.StatusInternalServerError)
		}
	}()

	// Parse request body
	var input models.UserRegisterInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		errors.ErrorResponseJSON(w, errors.ErrInvalidRequestPayload, http.StatusBadRequest)
		return
	}

	// Validate input
	err = validators.ValidateUserRegisterInput(input.Email, input.Password)
	if err != nil {
		errors.ErrorResponseJSON(w, errors.ErrInvalidInput, http.StatusBadRequest)
		return
	}

	// Check if the email already exists
	if emailExists(tx, input.Email) {
		errors.ErrorResponseJSON(w, errors.ErrEmailExists, http.StatusConflict)
		return
	}

	// Check if the phone number already exists
	if phoneNumberExists(tx, input.PhoneNumber) {
		errors.ErrorResponseJSON(w, errors.ErrPhoneNumberExists, http.StatusConflict)
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
		errors.ErrorResponseJSON(w, errors.ErrFailedToSetPassword, http.StatusInternalServerError)
		return
	}

	// Save the user to the appropriate database based on the config
	if cfg.MultipleDatabasesConfig.UseSQLite {
		err = tx.Create(user).Error
		if err != nil {
			tx.Rollback()
			errors.ErrorResponseJSON(w, errors.ErrFailedToSaveUserSQLite, http.StatusInternalServerError)
			return
		}
	} else if cfg.MultipleDatabasesConfig.UsePostgreSQL {
		err = tx.Create(user).Error
		if err != nil {
			tx.Rollback()
			errors.ErrorResponseJSON(w, errors.ErrFailedToSaveUserPostgreSQL, http.StatusInternalServerError)
			return
		}
		// TODO: Insert into PostgreSQL database
	} else {
		errors.ErrorResponseJSON(w, errors.ErrNoValidDatabaseSelected, http.StatusBadRequest)
		return
	}

	// Initialize Firebase app if configured to use Firebase
	if cfg.MultipleDatabasesConfig.UseRealtimeDatabase || cfg.MultipleDatabasesConfig.UseFirestore {
		ctx := context.Background()
		// Get a Firebase database client
		client, err := firebaseClient.Database(ctx)
		if err != nil {
			tx.Rollback()
			errors.ErrorResponseJSON(w, errors.ErrFailedToGetFirebaseClient, http.StatusInternalServerError)
			return
		}

		hashedPassword, _ := models.HashPassword(input.Password) // Hash the input password

		// Convert the user struct to a map
		userMap := map[string]interface{}{
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

		if cfg.MultipleDatabasesConfig.UseRealtimeDatabase {
			// Push the user data to Firebase Realtime Database
			ref := client.NewRef("users")
			newUserRef, err := ref.Push(context.Background(), userMap)
			if err != nil {
				tx.Rollback()
				errors.ErrorResponseJSON(w, errors.ErrFailedToSaveUserFirebaseRTDB, http.StatusInternalServerError)
				return
			}

			// Get the unique key generated by Firebase
			firebaseKey := newUserRef.Key
			log.Printf("Firebase Key: %s", firebaseKey)
		}

		if cfg.MultipleDatabasesConfig.UseFirestore {
			// TODO: Implement Firestore insertion
		}
	}

	// Commit the transaction
	tx.Commit()

	// Respond with success message or user data
	// For simplicity, sending a success response
	successResponse := models.UserResponse{
		Message: "User registered successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(successResponse)
}

// Function to check if email already exists in the database
func emailExists(db *gorm.DB, email string) bool {
	var user models.User
	err := db.Where("email = ?", email).First(&user).Error
	return err == nil
}

// Function to check if phone number already exists in the database
func phoneNumberExists(db *gorm.DB, phoneNumber string) bool {
	var user models.User
	err := db.Where("phone_number = ?", phoneNumber).First(&user).Error
	return err == nil
}
