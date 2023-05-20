package apperrors

// Generic errors
var (
	ServerError = "SERVER_ERROR"
)

// Invalid data received from the client
var (
	InvalidAuthMethod = "INVALID_AUTH_METHOD"
	InvalidEmail      = "INVALID_EMAIL"
	InvalidUsername   = "INVALID_USERNAME"
)

// Database errors
var (
	UsernameAlreadyExists = "USERNAME_ALREADY_EXISTS"
	EmailAlreadyExists    = "EMAIL_ALREADY_EXISTS"
)
