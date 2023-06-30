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

// NewAuthorizationCode creates a new authorization code
func NewAuthorizationCode(userId types.Id) (AuthorizationCodeModel, error) {
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

// PopulateAuthorizationCode creates a new authorization code and hydrates it with the given data
func PopulateAuthorizationCode(
	userId string,
	code string,
	createdAt time.Time,
	expiresAt time.Time) AuthorizationCodeModel {
	return AuthorizationCodeModel{
		UserId:    types.Id(userId),
		Code:      types.Code(code),
		CreatedAt: types.Timestamp(createdAt),
		ExpiresAt: types.Timestamp(expiresAt),
	}
}

func (v *AuthorizationCodeModel) IsExpired() bool {
	timeNow := types.NewTimestamp()
	return v.ExpiresAt.IsBefore(timeNow)
}
