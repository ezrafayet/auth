package model

import (
	"errors"
	"fmt"
	"iam/src/core/domain/types"
	"time"
)

type EmailVerificationCodeModel struct {
	UserId    types.Id
	Code      types.Code
	CreatedAt types.Timestamp
	ExpiresAt types.Timestamp
}

func NewVerificationCodeModel(userId types.Id) (EmailVerificationCodeModel, error) {
	timestamp := types.NewTimestamp()

	code, err := types.NewCode()

	if err != nil {
		fmt.Println(err)
		return EmailVerificationCodeModel{}, errors.New("SERVER_ERROR")
	}

	return EmailVerificationCodeModel{
		UserId:    userId,
		Code:      code,
		CreatedAt: timestamp,
		ExpiresAt: timestamp.AddSeconds(5 * 60),
	}, nil
}

func (v *EmailVerificationCodeModel) Hydrate(userId string, code string, createdAt time.Time, expiresAt time.Time) {
	v.UserId = types.Id(userId)
	v.Code = types.Code(code)
	v.CreatedAt = types.Timestamp(createdAt)
	v.ExpiresAt = types.Timestamp(expiresAt)
}

func (v *EmailVerificationCodeModel) IsExpired() bool {
	timeNow := types.NewTimestamp()
	return v.ExpiresAt.IsBefore(timeNow)
}
