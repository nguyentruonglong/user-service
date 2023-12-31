// User Model

package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents a user in the system.
type User struct {
	gorm.Model
	Email                       string     `gorm:"primaryKey;unique;not null" json:"email"`
	FirstName                   string     `json:"first_name" gorm:"not null"`
	MiddleName                  string     `json:"middle_name"`
	LastName                    string     `json:"last_name" gorm:"not null"`
	PasswordHash                string     `json:"password_hash"`
	IsActive                    bool       `json:"is_active" gorm:"default:true"`
	EmailVerificationCode       string     `json:"email_verification_code"`
	PhoneNumberVerificationCode string     `json:"phone_number_verification_code"`
	DateOfBirth                 time.Time  `json:"date_of_birth"`
	PhoneNumber                 string     `json:"phone_number"`
	Address                     string     `json:"address"`
	IsEmailVerified             bool       `json:"is_email_verified" gorm:"default:false"`
	IsPhoneNumberVerified       bool       `json:"is_phone_number_verified" gorm:"default:false"`
	Country                     string     `json:"country"`
	Province                    string     `json:"province"`
	AvatarURL                   string     `json:"avatar_url"`
	EarnedPoints                int        `json:"earned_points" gorm:"default:0"`
	ExtraInfo                   string     `json:"extra_info" gorm:"type:jsonb;default:'{}'"`
	CreatedAt                   time.Time  `json:"created_at"`
	UpdatedAt                   time.Time  `json:"updated_at"`
	DeletedAt                   *time.Time `gorm:"index"`
	Roles                       []Role     `gorm:"many2many:user_roles;default:null"`  // Define a many-to-many relationship with roles
	Groups                      []Group    `gorm:"many2many:user_groups;default:null"` // Define a many-to-many relationship with groups
}

// UserRegisterInput represents the input for user registration.
type UserRegisterInput struct {
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	MiddleName  string    `json:"middle_name"`
	LastName    string    `json:"last_name"`
	Password    string    `json:"password"`
	DateOfBirth time.Time `json:"date_of_birth"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Country     string    `json:"country"`
	Province    string    `json:"province"`
	AvatarURL   string    `json:"avatar_url"`
}

// UserLoginInput represents the input for user login.
type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserResetPasswordInput represents the input for resetting the user's password.
type UserResetPasswordInput struct {
	Email string `json:"email"`
}

// UserResponse represents the response for user-related operations.
type UserResponse struct {
	// Customize the response fields as needed
	Message string `json:"message"`
}

// TokenResponse represents the response containing a token.
type TokenResponse struct {
	Token string `json:"token"`
}

// LogoutResponse represents the response for a successful logout.
type LogoutResponse struct {
	Message string `json:"message"`
}

// BeforeCreate hook to set CreatedAt and UpdatedAt
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate hook to update UpdatedAt
func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	user.UpdatedAt = time.Now()
	return nil
}

// CreateVerificationCode creates a verification code for the user.
func (u *User) CreateVerificationCodes() {
	// Generate email verification code (e.g., a random six-digit code)
	u.EmailVerificationCode = generateVerificationCode()

	// Generate phone number verification code (e.g., a random six-digit code)
	u.PhoneNumberVerificationCode = generateVerificationCode()
}

func generateVerificationCode() string {
	// Implement logic to generate a verification code
	return "123456"
}

// HashPassword hashes the given password and returns the hashed password as a string.
func HashPassword(password string) (string, error) {
	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// SetPassword sets the hashed password for the user.
func (u *User) SetPassword(password string) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}
	u.PasswordHash = hashedPassword
	return nil
}

// IsValidPassword checks if the provided password matches the user's password.
func (u *User) IsValidPassword(password string) bool {
	// Compare the provided password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
