// Database Connection

package database

import (
	"log"
	"user-service/api/models"

	"gorm.io/driver/sqlite" // SQLite driver
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB initializes the SQLite database.
func InitDB(databaseURL string) (*gorm.DB, error) {
	// Open the SQLite database
	db, err := gorm.Open(sqlite.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Enable detailed log mode during development.
	db.Logger.LogMode(logger.Info)

	return db, nil
}

func AutoMigrateTables(db *gorm.DB) {
	// Migrate the User model
	err := db.AutoMigrate(&models.User{})
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
