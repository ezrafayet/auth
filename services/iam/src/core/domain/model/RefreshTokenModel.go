package model

import (
	"iam/src/core/domain/types"
	"time"
)

// RefreshTokenModel represents a refresh token
type RefreshTokenModel struct {
	UserId    types.Id
	CreatedAt types.Timestamp
	ExpiresAt types.Timestamp
	Revoked   bool
	RevokedAt types.Timestamp
	Token     types.Code
}

// NewRefreshToken creates a new refresh token
func NewRefreshToken(userId types.Id) (RefreshTokenModel, error) {
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

// PopulateRefreshToken creates a new refresh token and hydrates it with the given data
func PopulateRefreshToken(
	userId string,
	createdAt time.Time,
	expiresAt time.Time,
	revoked bool,
	revokedAt time.Time,
	token string) RefreshTokenModel {
	return RefreshTokenModel{
		UserId:    types.Id(userId),
		CreatedAt: types.Timestamp(createdAt),
		ExpiresAt: types.Timestamp(expiresAt),
		Token:     types.Code(token),
		Revoked:   revoked,
		RevokedAt: types.Timestamp(revokedAt),
	}
}

func (r *RefreshTokenModel) IsExpired() bool {
	return r.ExpiresAt.IsBefore(types.NewTimestamp())
}

func (r *RefreshTokenModel) Revoke() {
	r.Revoked = true
	r.RevokedAt = types.NewTimestamp()
}
