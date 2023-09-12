// Database Connection

package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // SQLite driver
)

// InitDB initializes the SQLite database.
func InitDB(databaseURL string) (*gorm.DB, error) {
	// Open the SQLite database
	db, err := gorm.Open("sqlite3", databaseURL)
	if err != nil {
		return nil, err
	}

	// Enable detailed log mode during development.
	db.LogMode(true)

	// Perform auto-migration to create/update database tables.
	autoMigrateTables(db)

	return db, nil
}

func autoMigrateTables(db *gorm.DB) {
	// Define database models and their relationships here.
	// Example:
	// db.AutoMigrate(&User{}, &Profile{}, &Post{})
	// You should define models in separate files.

	// Replace the following with actual models.
	// db.AutoMigrate(&YourModel{})
}

// CloseDB closes the database connection.
func CloseDB(db *gorm.DB) {
	if err := db.Close(); err != nil {
		log.Printf("Error closing the database: %v", err)
	}
}
