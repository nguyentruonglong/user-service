package models

import "context"

// ContextWithUserID creates a new context with the provided user ID.
func ContextWithUserID(ctx context.Context, userID uint) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// GetUserIDFromContext retrieves the user ID from the context.
func GetUserIDFromContext(ctx context.Context) (uint, bool) {
	userID, ok := ctx.Value(userIDKey).(uint)
	return userID, ok
}

// Key type for context value
type key int

const userIDKey key = iota
