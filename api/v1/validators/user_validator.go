package validators

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidInput       = errors.New("invalid input provided")
	ErrInvalidEmail       = errors.New("invalid email format")
	ErrWeakPassword       = errors.New("weak password")
	ErrInvalidPhoneNumber = errors.New("invalid phone number format")
)

// ValidateUserRegisterInput validates the user registration input.
func ValidateUserRegisterInput(email, password string) error {
	// Check if email and password are provided
	if email == "" || password == "" {
		return ErrInvalidInput
	}

	// Validate email format
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return ErrInvalidEmail
	}

	// Implement password strength validation
	if len(password) < 8 {
		return ErrWeakPassword
	}

	// Add more validation rules as needed

	return nil
}

// ValidateUserLoginInput validates the user login input.
func ValidateUserLoginInput(email, password string) error {
	// Check if email and password are provided
	if email == "" || password == "" {
		return ErrInvalidInput
	}

	// Validate email format
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return ErrInvalidEmail
	}

	// Add more validation rules as needed

	return nil
}

// ValidateUserResetPasswordInput validates the user reset password input.
func ValidateUserResetPasswordInput(email string) error {
	// Check if email is provided
	if email == "" {
		return ErrInvalidInput
	}

	// Validate email format
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return ErrInvalidEmail
	}

	// Add more validation rules as needed

	return nil
}
