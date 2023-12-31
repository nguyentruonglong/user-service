package models

import (
	"time"
)

// Role represents a role in the system and can have various permissions.
type Role struct {
	ID          uint         `gorm:"primaryKey"`
	Code        string       `gorm:"unique;not null"`
	Name        string       `gorm:"not null"`
	Description string       `gorm:"not null"`
	IsActive    bool         `gorm:"default:true"`
	CreatedAt   time.Time    `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time    `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time   `gorm:"index"`
	Permissions []Permission `gorm:"many2many:role_permissions;default:null"` // Define a many-to-many relationship with permissions
}
