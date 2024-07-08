package database

import (
	"log"
	"user-service/api/models"

	"gorm.io/driver/postgres" // PostgreSQL driver
	"gorm.io/driver/sqlite"   // SQLite driver
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitSQLiteDB initializes the SQLite database.
func InitSQLiteDB(connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Enable detailed log mode during development.
	db.Logger.LogMode(logger.Info)

	// Perform auto-migration of tables
	AutoMigrateTables(db)

	return db, nil
}

// InitPostgreSQLDB initializes the PostgreSQL database.
func InitPostgreSQLDB(connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Enable detailed log mode during development.
	db.Logger.LogMode(logger.Info)

	// Perform auto-migration of tables
	AutoMigrateTables(db)

	return db, nil
}

// AutoMigrateTables auto-migrates the database tables.
func AutoMigrateTables(db *gorm.DB) {
	// Migrate the User, Group, Permission, Role, AccessToken, RefreshToken, and EmailTemplate models
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
