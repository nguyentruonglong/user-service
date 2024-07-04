package models

import "context"

// Key type for context value to avoid conflicts
type key int

const (
	// userIDKey is the context key for the user ID
	userIDKey key = iota
)

// ContextWithUserID creates a new context with the provided user ID.
// It returns a new context that carries the user ID value.
func ContextWithUserID(ctx context.Context, userID uint) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// GetUserIDFromContext retrieves the user ID from the context.
// It returns the user ID and a boolean indicating whether the user ID was found in the context.
func GetUserIDFromContext(ctx context.Context) (uint, bool) {
	userID, ok := ctx.Value(userIDKey).(uint)
	return userID, ok
}
