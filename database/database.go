package database

import (
	"log"
	"user-service/api/models"
	"user-service/config"

	"gorm.io/driver/sqlite" // SQLite driver
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB initializes the SQLite database.
func InitDB(cfg *config.AppConfig) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// Get the multiple databases configuration
	dbConfig := cfg.GetMultipleDatabaseConfig()

	if dbConfig.GetUseSQLite() {
		db, err = gorm.Open(sqlite.Open(cfg.GetDatabaseURL()), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	}

	// Additional logic for other database configurations can be added here
	// e.g., PostgreSQL, Firebase, etc.

	// Enable detailed log mode during development.
	db.Logger.LogMode(logger.Info)

	// Perform auto-migration of tables
	AutoMigrateTables(db)

	return db, nil
}

// AutoMigrateTables auto-migrates the database tables.
func AutoMigrateTables(db *gorm.DB) {
	// Migrate the User and Token models
	err := db.AutoMigrate(
		&models.User{},
		&models.Group{},
		&models.Permission{},
		&models.Role{},
		&models.AccessToken{},
		&models.RefreshToken{},
		&models.EmailTemplate{},
	)

	if err != nil {
		log.Fatalf("Error auto-migrating tables: %v", err)
	}
}

// CloseDB closes the database connection.
func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error getting underlying SQL DB: %v", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing the database: %v", err)
	}
}
