package database

import (
	"context"
	"log"
	"os"
	"user-service/api/models"
	"user-service/config"
	"user-service/utils"

	firebase "firebase.google.com/go"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SeedEmailTemplates seeds initial email templates to the database and optionally to Firebase Realtime Database.
func SeedEmailTemplates(db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) error {
	// Read email template files
	emailVerificationBody, err := os.ReadFile("email_templates/email_verification.html")
	if err != nil {
		return err
	}

	passwordResetBody, err := os.ReadFile("email_templates/password_reset.html")
	if err != nil {
		return err
	}

	// Define email templates with detailed content and internal CSS
	templates := []models.EmailTemplate{
		{
			Code:        "EMAIL_VERIFICATION",
			Name:        "Email Verification",
			Subject:     "Verify your email address",
			Body:        string(emailVerificationBody),
			Params:      utils.ToJSONString([]string{"FirstName", "VerificationCode", "ExpiryTime"}),
			Description: "Template for verifying a user's email address",
		},
		{
			Code:        "PASSWORD_RESET",
			Name:        "Password Reset",
			Subject:     "Reset your password",
			Body:        string(passwordResetBody),
			Params:      utils.ToJSONString([]string{"FirstName", "ResetLink", "ExpiryTime"}),
			Description: "Template for resetting a user's password",
		},
	}

	// Check if SQLite or PostgreSQL is enabled and perform bulk upsert
	if cfg.DatabaseConfig.SQLite.Enabled || cfg.DatabaseConfig.PostgreSQL.Enabled {
		err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "code"}},
			UpdateAll: true,
		}).Create(&templates).Error
		if err != nil {
			return err
		}
	}

	// Check if Firebase Realtime Database is enabled and perform operations
	if cfg.DatabaseConfig.Firebase.Enabled {
		ctx := context.Background()

		// Get Firebase database client
		client, err := firebaseClient.Database(ctx)
		if err != nil {
			return err
		}

		// Fetch existing templates from Firebase
		var existingTemplates map[string]models.EmailTemplate
		ref := client.NewRef("email_templates")
		if err := ref.Get(ctx, &existingTemplates); err != nil {
			log.Printf("Error fetching data from Firebase: %v", err)
			return err
		}

		// Initialize map if nil
		if existingTemplates == nil {
			existingTemplates = make(map[string]models.EmailTemplate)
		}

		// Merge or update the templates
		for _, template := range templates {
			existingTemplates[template.Code] = template
		}

		// Save the merged templates back to Firebase
		if err := ref.Set(ctx, existingTemplates); err != nil {
			log.Printf("Error saving data to Firebase: %v", err)
			return err
		}
	}

	return nil
}
