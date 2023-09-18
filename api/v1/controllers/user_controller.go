// User Controller

package controllers

import (
	"encoding/json"
	"net/http"
	"user-service/api/models"
	"user-service/api/v1/validators"

	"gorm.io/gorm"
)

// @Summary Register a new user
// @Description Register a new user with the given email and password.
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} models.UserResponse
// @Failure 409 {object} models.UserResponse
// @Failure 500 {object} models.UserResponse
// @Router /api/v1/register [post]
func RegisterUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
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
	if emailExists(db, input.Email) {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	// Check if the phone number already exists
	if phoneNumberExists(db, input.PhoneNumber) {
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
		http.Error(w, "Failed to set the password", http.StatusInternalServerError)
		return
	}

	// Save the user to the database
	err = db.Create(user).Error

	if err != nil {
		http.Error(w, "Failed to save user to the database", http.StatusInternalServerError)
		return
	}

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
