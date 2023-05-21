package primaryports

import (
	"iam/src/core/domain/types"
)

// Interfaces for service: user

type RegisterArgs struct {
	AuthMethod string `json:"authMethod"`
	Email      string `json:"email"`
	Username   string `json:"username"`
}

type RegisterAnswer struct {
	UserId types.Id `json:"userId"`
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
