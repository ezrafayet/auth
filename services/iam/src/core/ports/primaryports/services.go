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
	Authenticate(args GetAccessTokenArgs) (GetAccessTokenAnswer, error)
}

type AuthorizationService interface {
	IsAccessTokenValid(args IsAccessTokenValidArgs) (IsAccessTokenValidAnswer, error)
	RefreshAccessToken(args RefreshAccessTokenArgs) (RefreshAccessTokenAnswer, error)
}
