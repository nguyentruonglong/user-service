package database

import (
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

	// if cfg.GetMultipleDatabasesConfig().GetUseRealtimeDatabase() {
	// 	ctx := context.Background()

	// 	// Fetch existing templates from Firebase
	// 	existingTemplates := make(map[string]models.EmailTemplate)
	// 	err := firebase_services.GetDataFromRealtimeDB(ctx, "email_templates", &existingTemplates)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// Merge or update the templates
	// 	for _, template := range templates {
	// 		existingTemplates[template.Code] = template
	// 	}

	// 	// Save the merged templates back to Firebase
	// 	path := "email_templates"
	// 	if err := firebase_services.SaveDataToRealtimeDB(ctx, path, existingTemplates); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}
