package models

import (
	"time"

	"gorm.io/gorm"
)

// Role represents a role in the system and can have various permissions.
type Role struct {
	ID          uint         `gorm:"primaryKey"`
	Code        string       `gorm:"unique;not null"`
	Name        string       `gorm:"not null"`
	Description string       `gorm:"not null"`
	IsActive    bool         `gorm:"default:true"`
	CreatedAt   time.Time    `gorm:"autoCreateTime"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime"`
	DeletedAt   *time.Time   `gorm:"index"`
	Permissions []Permission `gorm:"many2many:role_permissions"` // Define a many-to-many relationship with permissions
}

// BeforeCreate hook to set default values
func (role *Role) BeforeCreate(tx *gorm.DB) (err error) {
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()
	return
}

// BeforeUpdate hook to update the UpdatedAt field
func (role *Role) BeforeUpdate(tx *gorm.DB) (err error) {
	role.UpdatedAt = time.Now()
	return
}
