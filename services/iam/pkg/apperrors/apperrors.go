package apperrors

// Environment variables not set
var (
	DatabaseConnectionStringNotSet = "DATABASE_CONNECTION_STRING_NOT_SET"
	EmailProviderNotSet            = "EMAIL_PROVIDER_NOT_SET"
)

// Generic errors
var (
	ServerError   = "SERVER_ERROR"
	LimitExceeded = "LIMIT_EXCEEDED"
)

// Invalid data received from the client
var (
	InvalidId           = "INVALID_ID"
	InvalidAuthType     = "INVALID_AUTH_TYPE"
	InvalidRole         = "INVALID_ROLE"
	InvalidEmail        = "INVALID_EMAIL"
	InvalidUsername     = "INVALID_USERNAME"
	InvalidCode         = "INVALID_CODE"
	RefusedTerms        = "REFUSED_TERMS"
	InvalidAccessToken  = "INVALID_ACCESS_TOKEN"
	InvalidRefreshToken = "INVALID_REFRESH_TOKEN"
	FeatureDisabled     = "FEATURE_DISABLED"
)

// Database errors
var (
	UsernameAlreadyExists = "USERNAME_ALREADY_EXISTS"
	EmailAlreadyExists    = "EMAIL_ALREADY_EXISTS"
	UserNotFound          = "USER_NOT_FOUND"
)

// Service errors
var (
	EmailNotVerified          = "EMAIL_NOT_VERIFIED"
	EmailAlreadyVerified      = "EMAIL_ALREADY_VERIFIED"
	VerificationCodeNotFound  = "VERIFICATION_NOT_FOUND"
	AuthorizationCodeNotFound = "AUTHORIZATION_CODE_NOT_FOUND"
	VerificationCodeExpired   = "VERIFICATION_CODE_EXPIRED"
	AuthorizationCodeExpired  = "AUTHORIZATION_CODE_EXPIRED"
	UserBlocked               = "USER_BLOCKED"
	UserDeleted               = "USER_DELETED"
)
