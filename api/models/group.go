package models

import (
	"time"

	"gorm.io/gorm"
)

// Group represents a group in the system and can have multiple users.
type Group struct {
	ID          uint       `gorm:"primaryKey"`
	Code        string     `gorm:"unique;not null"`
	Name        string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	IsActive    bool       `gorm:"default:true"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `gorm:"index"`
	Users       []User     `gorm:"many2many:user_groups"` // Define a many-to-many relationship with users
}

// BeforeCreate hook to set default values
func (group *Group) BeforeCreate(tx *gorm.DB) (err error) {
	group.CreatedAt = time.Now()
	group.UpdatedAt = time.Now()
	return
}

// BeforeUpdate hook to update the UpdatedAt field
func (group *Group) BeforeUpdate(tx *gorm.DB) (err error) {
	group.UpdatedAt = time.Now()
	return
}
