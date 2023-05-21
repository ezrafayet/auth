package model

import (
	"fmt"
	"iam/src/core/domain/types"
	"time"
)

type AuthorizationCodeModel struct {
	UserId    types.Id
	Code      types.Code
	CreatedAt types.Timestamp
	ExpiresAt types.Timestamp
}

func NewAuthorizationCodeModel(userId types.Id) (AuthorizationCodeModel, error) {
	timestamp := types.NewTimestamp()

	code, err := types.NewCode()

	if err != nil {
		fmt.Println(err)
		return AuthorizationCodeModel{}, err
	}

	return AuthorizationCodeModel{
		UserId:    userId,
		Code:      code,
		CreatedAt: timestamp,
		ExpiresAt: timestamp.AddSeconds(60),
	}, nil
}

func (v *AuthorizationCodeModel) Hydrate(userId string, code string, createdAt time.Time, expiresAt time.Time) {
	v.UserId = types.Id(userId)
	v.Code = types.Code(code)
	v.CreatedAt = types.Timestamp(createdAt)
	v.ExpiresAt = types.Timestamp(expiresAt)
}

func (v *AuthorizationCodeModel) IsExpired() bool {
	timeNow := types.NewTimestamp()
	return v.ExpiresAt.IsBefore(timeNow)
}
