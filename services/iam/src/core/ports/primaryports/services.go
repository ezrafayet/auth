package primaryports

type UsersService interface {
	Register(args RegisterArgs) (RegisterAnswer, error)
}

type EmailVerificationService interface {
	Send(args SendVerificationCodeArgs) error
	Confirm(args ConfirmEmailArgs) (ConfirmEmailAnswer, error)
}

type AuthenticationService interface {
	SendMagicLink(args SendMagicLinkArgs) (SendMagicLinkAnswer, error)
}
