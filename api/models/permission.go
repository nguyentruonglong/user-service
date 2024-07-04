package models

import (
	"time"

	"gorm.io/gorm"
)

// Permission represents a permission in the system.
type Permission struct {
	ID          uint      `gorm:"primaryKey"`
	Code        string    `gorm:"unique;not null"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	IsActive    bool      `gorm:"default:true"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

// BeforeCreate hook to set default values
func (permission *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	permission.CreatedAt = time.Now()
	permission.UpdatedAt = time.Now()
	return
}

// BeforeUpdate hook to update the UpdatedAt field
func (permission *Permission) BeforeUpdate(tx *gorm.DB) (err error) {
	permission.UpdatedAt = time.Now()
	return
}
