// Token Model

package models

import (
	"time"

	"gorm.io/gorm"
)

// AccessToken represents a JWT token stored in the database instead of caching.
// This is not a good idea, the access token should be stored in cache.
type AccessToken struct {
	gorm.Model
	UserID         uint      `gorm:"not null;index:idx_user_access_token_expiration_time"` // User ID associated with the token
	AccessToken    string    `gorm:"uniqueIndex;not null;index:idx_user_access_token_expiration_time"`
	ExpirationTime time.Time `gorm:"index:idx_user_access_token_expiration_time"`
}

// RefreshToken represents a refresh token stored in the database.
type RefreshToken struct {
	gorm.Model
	UserID         uint      `gorm:"not null;index:idx_user_refresh_token_expiration_time"` // User ID associated with the token
	RefreshToken   string    `gorm:"uniqueIndex;not null;index:idx_user_refresh_token_expiration_time"`
	ExpirationTime time.Time `gorm:"index:idx_user_refresh_token_expiration_time"`
}
