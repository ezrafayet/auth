package primaryports

import "iam/src/core/services"

type UsersService interface {
	Register(args services.RegisterArgs) (services.RegisterAnswer, error)
	// WhoAmI() error
}

type EmailVerificationService interface {
	Send(args services.SendVerificationCodeArgs) error
	Verify(args services.VerifyCodeArgs) (services.VerifyCodeAnswer, error)
}

type AuthenticationService interface {
	// MagicLink() error
	// Token() error
	// RefreshToken() error
	// RevokeToken() error
	// VerifyToken() error
}
