package models

import (
	"time"
)

// Permission represents a permission in the system.
type Permission struct {
	ID          uint      `gorm:"primaryKey"`
	Code        string    `gorm:"unique;not null"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	IsActive    bool      `gorm:"default:true"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
