package errors

import "errors"

// Common error messages
var (
	// ErrTransactionFailed occurs when a transaction fails.
	ErrTransactionFailed = errors.New("transaction failed")

	// ErrInvalidRequestPayload occurs when the request payload is invalid.
	ErrInvalidRequestPayload = errors.New("invalid request payload")

	// ErrInvalidInput occurs when the input provided is invalid.
	ErrInvalidInput = errors.New("invalid input")

	// ErrEmailExists occurs when the email already exists.
	ErrEmailExists = errors.New("email already exists")

	// ErrPhoneNumberExists occurs when the phone number already exists.
	ErrPhoneNumberExists = errors.New("phone number already exists")

	// ErrFailedToSetPassword occurs when setting the password fails.
	ErrFailedToSetPassword = errors.New("failed to set the password")

	// ErrFailedToSaveUserSQLite occurs when saving a user to SQLite database fails.
	ErrFailedToSaveUserSQLite = errors.New("failed to save user to the SQLite database")

	// ErrFailedToSaveUserPostgreSQL occurs when saving a user to PostgreSQL database fails.
	ErrFailedToSaveUserPostgreSQL = errors.New("failed to save user to the PostgreSQL database")

	// ErrNoValidDatabaseSelected occurs when no valid database is selected.
	ErrNoValidDatabaseSelected = errors.New("no valid database selected")

	// ErrFailedToGetFirebaseClient occurs when getting the Firebase database client fails.
	ErrFailedToGetFirebaseClient = errors.New("failed to get Firebase database client")

	// ErrFailedToSaveUserFirebaseRTDB occurs when saving a user to Firebase Realtime Database fails.
	ErrFailedToSaveUserFirebaseRTDB = errors.New("failed to save user to Firebase Realtime Database")
)
