package database

import (
	"context"
	"log"
	"user-service/api/models"
	"user-service/config"
	"user-service/utils"

	firebase "firebase.google.com/go"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SeedEmailTemplates seeds initial email templates to the database and optionally to Firebase Realtime Database.
func SeedEmailTemplates(db *gorm.DB, firebaseClient *firebase.App, cfg *config.AppConfig) error {
	templates := []models.EmailTemplate{
		{
			Code:        "EMAIL_VERIFICATION",
			Name:        "Email Verification",
			Subject:     "Verify your email address",
			Body:        "<p>Dear {{.FirstName}},</p><p>Please click the link below to verify your email address:</p><p><a href=\"{{.VerificationLink}}\">Verify Email</a></p><p>Thank you!</p>",
			Params:      utils.ToJSONString([]string{"FirstName", "VerificationLink"}),
			Description: "Template for verifying a user's email address",
		},
		{
			Code:        "PASSWORD_RESET",
			Name:        "Password Reset",
			Subject:     "Reset your password",
			Body:        "<p>Dear {{.FirstName}},</p><p>Please click the link below to reset your password:</p><p><a href=\"{{.ResetLink}}\">Reset Password</a></p><p>If you did not request a password reset, please ignore this email.</p>",
			Params:      utils.ToJSONString([]string{"FirstName", "ResetLink"}),
			Description: "Template for resetting a user's password",
		},
	}

	if cfg.GetMultipleDatabasesConfig().GetUseSQLite() || cfg.GetMultipleDatabasesConfig().GetUsePostgreSQL() {
		// Bulk upsert operation for SQLite or PostgreSQL
		err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "code"}},
			UpdateAll: true,
		}).Create(&templates).Error
		if err != nil {
			return err
		}
	}

	if cfg.GetMultipleDatabasesConfig().GetUseRealtimeDatabase() {
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
