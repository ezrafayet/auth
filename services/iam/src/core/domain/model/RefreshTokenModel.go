package model

import (
	"iam/src/core/domain/types"
)

type RefreshTokenModel struct {
	UserId    types.Id
	CreatedAt types.Timestamp
	ExpiresAt types.Timestamp
	Revoked   bool
	RevokedAt types.Timestamp
	Token     types.Code
}

func NewRefreshTokenModel(userId types.Id) (RefreshTokenModel, error) {
	timestamp := types.NewTimestamp()

	token, err := types.NewCode()

	if err != nil {
		return RefreshTokenModel{}, err
	}

	return RefreshTokenModel{
		UserId:    userId,
		CreatedAt: timestamp,
		ExpiresAt: timestamp.AddMonths(3),
		Token:     token,
	}, nil
}
