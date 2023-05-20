package primaryports

type UsersService interface {
	Register(args RegisterArgs) (RegisterAnswer, error)
	// WhoAmI() error
}

type EmailVerificationService interface {
	Send(args SendVerificationCodeArgs) error
	Confirm(args ConfirmEmailArgs) (ConfirmEmailAnswer, error)
}

type AuthenticationService interface {
	// MagicLink() error
	// Token() error
	// RefreshToken() error
	// RevokeToken() error
	// VerifyToken() error
}
