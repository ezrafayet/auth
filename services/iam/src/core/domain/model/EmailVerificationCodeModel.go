package model

import (
	"iam/src/core/domain/types"
	"time"
)

// EmailVerificationCodeModel represents a verification code for an email address
type EmailVerificationCodeModel struct {
	UserId    types.Id
	Code      types.Code
	CreatedAt types.Timestamp
	ExpiresAt types.Timestamp
}

// NewEmailVerificationCode creates a new email verification code
func NewEmailVerificationCode(userId types.Id) (EmailVerificationCodeModel, error) {
	timestamp := types.NewTimestamp()

	code, err := types.NewCode()

	if err != nil {
		return EmailVerificationCodeModel{}, err
	}

	return EmailVerificationCodeModel{
		UserId:    userId,
		Code:      code,
		CreatedAt: timestamp,
		ExpiresAt: timestamp.AddSeconds(5 * 60),
	}, nil
}

// PopulateEmailVerificationCode creates a new email verification code and hydrates it with the given data
func PopulateEmailVerificationCode(
	userId string,
	code string,
	createdAt time.Time,
	expiresAt time.Time) EmailVerificationCodeModel {
	return EmailVerificationCodeModel{
		UserId:    types.Id(userId),
		Code:      types.Code(code),
		CreatedAt: types.Timestamp(createdAt),
		ExpiresAt: types.Timestamp(expiresAt),
	}
}

func (v *EmailVerificationCodeModel) IsExpired() bool {
	return v.ExpiresAt.IsBefore(types.NewTimestamp())
}
