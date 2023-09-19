// User Controller

package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"user-service/api/models"
	"user-service/api/v1/validators"
	"user-service/config"
	"user-service/utils"

	firebase "firebase.google.com/go"

	"gorm.io/gorm"
)

// @Summary Register a new user
// @Description Register a new user with the given email and password.
// @Accept json
// @Produce json
// @Param email body string true "Email (required)"
// @Param password body string true "Password (required)"
// @Param first_name body string false "First Name (required)"
// @Param middle_name body string false "Middle Name (optional)"
// @Param last_name body string false "Last Name (required)"
// @Param date_of_birth body string false "Date of Birth (optional) (YYYY-MM-DD)"
// @Param phone_number body string false "Phone Number (optional)"
// @Param address body string false "Address (optional)"
// @Param country body string false "Country (optional)"
// @Param province body string false "Province (optional)"
// @Param avatar_url body string false "Avatar URL (optional)"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} models.UserResponse
// @Failure 409 {object} models.UserResponse
// @Failure 500 {object} models.UserResponse
// @Router /api/v1/register [post]

// RegisterUser registers a new user and creates an object in the specified database
func RegisterUser(w http.ResponseWriter, r *http.Request, db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) {
	// Start a transaction
	tx := db.Begin()

	// Rollback the transaction if any error occurs
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			http.Error(w, "Transaction failed", http.StatusInternalServerError)
		}
	}()

	// Parse request body
	var input models.UserRegisterInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate input
	err = validators.ValidateUserRegisterInput(input.Email, input.Password)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the email already exists
	if emailExists(tx, input.Email) {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	// Check if the phone number already exists
	if phoneNumberExists(tx, input.PhoneNumber) {
		http.Error(w, "Phone number already exists", http.StatusConflict)
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
		http.Error(w, "Failed to set the password", http.StatusInternalServerError)
		return
	}

	// Save the user to the appropriate database based on the config
	if cfg.MultipleDatabasesConfig.UseSQLite {
		err = tx.Create(user).Error
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to save user to the SQLite database", http.StatusInternalServerError)
			return
		}
	} else if cfg.MultipleDatabasesConfig.UsePostgreSQL {
		err = tx.Create(user).Error
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to save user to the PostgreSQL database", http.StatusInternalServerError)
			return
		}
		// TODO: Insert into PostgreSQL database
	} else {
		http.Error(w, "No valid database selected", http.StatusBadRequest)
		return
	}

	// Initialize Firebase app if configured to use Firebase
	if cfg.MultipleDatabasesConfig.UseRealtimeDatabase || cfg.MultipleDatabasesConfig.UseFirestore {
		ctx := context.Background()
		// Get a Firebase database client
		client, err := firebaseClient.Database(ctx)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to get Firebase database client", http.StatusInternalServerError)
			return
		}

		hashedPassword, _ := models.HashPassword(input.Password) // Hash input password

		// Convert the user struct to a map
		userMap := map[string]interface{}{
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
			"password_hash":                  hashedPassword,                     // Add default password hash value
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
				http.Error(w, "Failed to save user to Firebase Realtime Database", http.StatusInternalServerError)
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
