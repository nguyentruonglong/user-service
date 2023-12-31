package models

import (
	"time"
)

// Group represents a group in the system and can have multiple users.
type Group struct {
	ID          uint       `gorm:"primaryKey"`
	Code        string     `gorm:"unique;not null"`
	Name        string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	IsActive    bool       `gorm:"default:true"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time `gorm:"index"`
	Users       []User     `gorm:"many2many:user_groups;default:null"` // Define a many-to-many relationship with users
}
