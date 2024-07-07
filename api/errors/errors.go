package errors

import "errors"

// Common error messages
var (
	// ErrTransactionFailed occurs when a transaction fails.
	ErrTransactionFailed = errors.New("transaction failed")

	// ErrInvalidRequestPayload occurs when the request payload is invalid.
	ErrInvalidRequestPayload = errors.New("invalid request payload")

	// ErrInvalidInput occurs when the input provided is invalid.
	ErrInvalidInput = errors.New("invalid input provided")

	// ErrInvalidEmail is returned for invalid email format provided.
	ErrInvalidEmail = errors.New("invalid email format provided")

	// ErrWeakPassword is returned for weak password provided.
	ErrWeakPassword = errors.New("weak password provided")

	// ErrInvalidPhoneNumber is returned for invalid phone number format provided.
	ErrInvalidPhoneNumber = errors.New("invalid phone number format provided")

	// ErrEmailExistsInDatabase indicates that the email already exists in the database.
	ErrEmailExistsInDatabase = errors.New("email already exists in the database")

	// ErrPhoneNumberExistsInDatabase indicates that the phone number already exists in the database.
	ErrPhoneNumberExistsInDatabase = errors.New("phone number already exists in the database")

	// ErrPhoneNumberNotFoundInDatabase indicates that the phone number does not exist in the database.
	ErrPhoneNumberNotFoundInDatabase = errors.New("phone number not found in the database")

	// ErrFailedToCheckPhoneNumberExistence indicates failure to check if a phone number exists.
	ErrFailedToCheckPhoneNumberExistence = errors.New("failed to check phone number existence")

	// ErrPhoneNumberAlreadyExistsOnFirebase indicates that the phone number already exists in Firebase.
	ErrPhoneNumberAlreadyExistsOnFirebase = errors.New("phone number already exists in Firebase")

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

	// ErrFailedToCheckEmailExistence is an error when checking email existence in Firebase.
	ErrFailedToCheckEmailExistence = errors.New("failed to check email existence")

	// ErrEmailAlreadyExistsOnFirebase is an error when the email already exists in Firebase.
	ErrEmailAlreadyExistsOnFirebase = errors.New("email already exists in Firebase")

	// ErrAuthenticationFailed is an error when user authentication fails.
	ErrAuthenticationFailed = errors.New("authentication failed")

	// ErrTokenGenerationFailed is an error when token generation fails.
	ErrTokenGenerationFailed = errors.New("token generation failed")

	// ErrUnauthorized occurs when a request is unauthorized.
	ErrUnauthorized = errors.New("authentication failed or insufficient permissions")

	// ErrDatabaseOperationFailed occurs when a database operation fails.
	ErrDatabaseOperationFailed = errors.New("database operation failed")

	// ErrInvalidToken is returned when the provided token is invalid.
	ErrInvalidToken = errors.New("invalid token")

	// ErrInvalidRefreshToken is returned when the provided refresh token is invalid or missing.
	ErrInvalidRefreshToken = errors.New("invalid or missing refresh token")

	// ErrUserNotFound occurs when a user is not found in the database.
	ErrUserNotFound = errors.New("user not found")

	// ErrSMSFailure occurs when sending an SMS fails.
	ErrSMSFailure = errors.New("SMS sending failure")

	// ErrInvalidEmailVerificationInput occurs when the email verification input is invalid.
	ErrInvalidEmailVerificationInput = errors.New("invalid email verification input")

	// ErrEmailTaskPublishingFailed occurs when publishing an email task to the queue fails.
	ErrEmailTaskPublishingFailed = errors.New("failed to publish email task")

	// ErrEmailAlreadyVerified indicates that the email is already verified.
	ErrEmailAlreadyVerified = errors.New("email already verified")

	// ErrEmailNotProvided indicates that the email is not provided.
	ErrEmailNotProvided = errors.New("email not provided")
)
