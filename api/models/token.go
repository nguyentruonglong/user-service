// Token Model

package models

import (
	"time"

	"gorm.io/gorm"
)

// Token represents a JWT token stored in the database.
type Token struct {
	gorm.Model
	UserID         uint      `gorm:"not null;index:idx_user_access_token_expiration_time"` // User ID associated with the token
	AccessToken    string    `gorm:"uniqueIndex;not null;index:idx_user_access_token_expiration_time"`
	ExpirationTime time.Time `gorm:"index:idx_user_access_token_expiration_time"`
}
