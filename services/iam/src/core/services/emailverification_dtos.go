package services

import (
	"iam/src/core/types"
)

type SendVerificationCodeArgs struct {
	UserId types.Id `json:"userId"`
}

type SendVerificationCodeAnswer struct{}

type VerifyCodeArgs struct {
	VerificationCode string `json:"verificationCode"`
}

type VerifyCodeAnswer struct {
	AuthorizationCode string `json:"authorizationCode"`
}
