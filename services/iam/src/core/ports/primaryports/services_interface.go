package primaryports

// Interfaces for service: user

type RegisterArgs struct {
	AuthType              string `json:"authType"`
	Email                 string `json:"email"`
	Username              string `json:"username"`
	HasAcceptedTerms      bool   `json:"hasAcceptedTerms"`
	AcceptedTermsVersion  string `json:"acceptedTermsVersion"`
	HasAcceptedNewsletter bool   `json:"hasAcceptedNewsletter"`
	HasAcceptedMarketing  bool   `json:"hasAcceptedMarketing"`
}

type RegisterAnswer struct {
	UserId string `json:"userId"`
}

// Interfaces for service: email verification

type SendVerificationCodeArgs struct {
	UserId string `json:"userId"`
}

type SendVerificationCodeAnswer struct{}

type ConfirmEmailArgs struct {
	VerificationCode string `json:"verificationCode"`
}

type ConfirmEmailAnswer struct {
	AuthorizationCode string `json:"authorizationCode"`
}

// Interfaces for service: authentication

type SendMagicLinkArgs struct {
	Email string `json:"email"`
}

type SendMagicLinkAnswer struct {
	AuthorizationCode string `json:"authorizationCode"`
}

type GetAccessTokenArgs struct {
	AuthorizationCode string `json:"authorizationCode"`
}

type GetAccessTokenAnswer struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// Interfaces for service: authorization

type IsAccessTokenValidArgs struct {
	AuthorisationHeader string
}

type IsAccessTokenValidAnswer struct {
	Valid bool `json:"valid"`
}
