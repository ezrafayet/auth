package model

import (
	"iam/src/core/types"
)

type EmailVerificationCodeModel struct {
	UserId    types.Id
	Code      types.Code
	CreatedAt types.Timestamp
	ExpiresAt types.Timestamp
}

func NewVerificationCodeModel(userId types.Id, code types.Code, timestamp types.Timestamp, expiresIn int) EmailVerificationCodeModel {
	return EmailVerificationCodeModel{
		UserId:    userId,
		Code:      code,
		CreatedAt: timestamp,
		ExpiresAt: timestamp.AddSeconds(expiresIn),
	}
}

func (v *EmailVerificationCodeModel) IsExpired() bool {
	timeNow := types.NewTimestamp()
	return v.ExpiresAt.IsBefore(timeNow)
}
