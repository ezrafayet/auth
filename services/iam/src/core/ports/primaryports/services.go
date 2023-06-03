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
	RefreshAccessToken(args RefreshAccessTokenArgs) (RefreshAccessTokenAnswer, error)
	IsAccessTokenValid(args IsAccessTokenValidArgs) (IsAccessTokenValidAnswer, error)
	AreFeaturesEnabled(args AreFeaturesEnabledArgs) (AreFeaturesEnabledAnswer, error)
}
