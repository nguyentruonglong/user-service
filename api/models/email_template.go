package models

import (
	"time"

	"gorm.io/gorm"
)

// EmailTemplate represents a template for outgoing emails.
type EmailTemplate struct {
	ID          uint      `gorm:"primaryKey"`
	Code        string    `gorm:"unique;not null"`    // Unique code for the template
	Name        string    `gorm:"not null"`           // Name of the template
	Subject     string    `gorm:"not null"`           // Subject of the email
	Body        string    `gorm:"type:text;not null"` // Body of the email
	Params      string    `gorm:"type:text;not null"` // Parameters for the template in JSON format
	Description string    `gorm:"not null"`           // Description of the template
	CreatedAt   time.Time `gorm:"autoCreateTime"`     // Time of creation
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`     // Time of last update
}

// BeforeCreate sets timestamps before creating a new record.
func (e *EmailTemplate) BeforeCreate(tx *gorm.DB) (err error) {
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return
}

// BeforeUpdate sets the timestamp before updating a record.
func (e *EmailTemplate) BeforeUpdate(tx *gorm.DB) (err error) {
	e.UpdatedAt = time.Now()
	return
}

// SendEmailVerificationCodeInput represents the input data for sending email verification.
type SendEmailVerificationCodeInput struct {
}

// SendEmailVerificationCodeResponse represents the response after sending the verification email.
type SendEmailVerificationCodeResponse struct {
	Message string `json:"message"`
}

// EmailVerificationInput represents the input data for verifying the email.
type EmailVerificationInput struct {
	VerificationCode string `json:"verification_code" binding:"required"`
}

// EmailVerificationResponse represents the response after verifying the email.
type EmailVerificationResponse struct {
	Message string `json:"message"`
}
