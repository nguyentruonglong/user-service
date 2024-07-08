// Authentication Middleware

package middlewares

import (
	"net/http"
	"strings"
	"time"
	"user-service/api/errors"
	"user-service/api/models"
	"user-service/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AuthMiddleware is a middleware function to validate the access token.
func AuthMiddleware(db *gorm.DB, cfg *config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			errors.ErrorResponseJSON(c.Writer, errors.ErrUnauthorized, http.StatusUnauthorized)
			c.Abort()
			return
		}

		// Extract the token from the "Bearer" prefix
		var accessToken string
		tokenParts := strings.Split(authorizationHeader, " ")
		if len(tokenParts) == 2 && tokenParts[0] == "Bearer" {
			accessToken = tokenParts[1]
		} else {
			// No "Bearer" prefix found, consider the whole value as the token
			accessToken = authorizationHeader
		}

		// Validate the token
		userID, err := validateToken(accessToken, db, cfg)
		if err != nil {
			errors.ErrorResponseJSON(c.Writer, err, http.StatusUnauthorized)
			c.Abort()
			return
		}

		// Add user ID to the request context for controllers to use
		c.Set("userID", userID)

		c.Next()
	}
}

// validateToken validates the access token and returns the user ID.
func validateToken(accessToken string, db *gorm.DB, cfg *config.AppConfig) (uint, error) {
	// Parse the token
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrInvalidToken
		}

		// Provide the secret key for validation
		return []byte(cfg.JWTSecretKey), nil
	})

	// Check for parsing errors and token validity
	if err != nil || !token.Valid {
		validationError, _ := err.(*jwt.ValidationError)
		if validationError.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return 0, errors.ErrInvalidToken
		}
		return 0, errors.ErrInvalidToken
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.ErrInvalidToken
	}

	// Extract user ID from claims
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.ErrInvalidToken
	}

	userID := uint(userIDFloat)

	// Check if the token exists in the database and is not expired
	var storedToken models.AccessToken
	err = db.Where("user_id = ? AND access_token = ? AND expiration_time > ?", userID, accessToken, time.Now()).First(&storedToken).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, errors.ErrInvalidToken
		}
		return 0, errors.ErrDatabaseOperationFailed
	}

	return userID, nil
}
