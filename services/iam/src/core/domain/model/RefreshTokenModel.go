package model

import (
	"iam/src/core/domain/types"
	"time"
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

func (r *RefreshTokenModel) Hydrate(
	userId string,
	createdAt time.Time,
	expiresAt time.Time,
	token string,
	revoked bool,
	revokedAt time.Time) error {
	r.UserId = types.Id(userId)
	r.CreatedAt = types.Timestamp(createdAt)
	r.ExpiresAt = types.Timestamp(expiresAt)
	r.Token = types.Code(token)
	r.Revoked = revoked
	r.RevokedAt = types.Timestamp(revokedAt)
	return nil
}

func (r *RefreshTokenModel) IsExpired() bool {
	return r.ExpiresAt.IsBefore(types.NewTimestamp())
}

func (r *RefreshTokenModel) Revoke() {
	r.Revoked = true
	r.RevokedAt = types.NewTimestamp()
}
