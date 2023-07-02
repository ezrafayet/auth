package model

import (
	"time"

	"iam/src/core/domain/types"
)

type UserModel struct {
	Id                  types.Id
	CreatedAt           types.Timestamp
	Username            types.Username
	UsernameFingerprint types.UsernameFingerprint
	Email               types.Email
	EmailVerified       bool
	EmailVerifiedAt     types.Timestamp
	Blocked             bool
	Deleted             bool
	DeletedAt           types.Timestamp
}

func NewUser(username types.Username, email types.Email, timestamp types.Timestamp) UserModel {
	return UserModel{
		Id:                  types.NewId(),
		CreatedAt:           timestamp,
		Username:            username,
		UsernameFingerprint: types.ComputeUsernameFingerprint(username),
		Email:               email,
		EmailVerified:       false,
		Blocked:             false,
		Deleted:             false,
	}
}

func PopulateUser(
	id string,
	createdAt time.Time,
	username string,
	usernameFingerprint string,
	email string,
	emailVerified bool,
	emailVerifiedAt time.Time,
	blocked bool,
	deleted bool,
	deletedAt time.Time) UserModel {
	return UserModel{
		Id:                  types.Id(id),
		CreatedAt:           types.Timestamp(createdAt),
		Username:            types.Username(username),
		UsernameFingerprint: types.UsernameFingerprint(usernameFingerprint),
		Email:               types.Email(email),
		EmailVerified:       emailVerified,
		EmailVerifiedAt:     types.Timestamp(emailVerifiedAt),
		Blocked:             blocked,
		Deleted:             deleted,
		DeletedAt:           types.Timestamp(deletedAt),
	}
}

func (u *UserModel) IsBlocked() bool {
	return u.Blocked
}

func (u *UserModel) IsDeleted() bool {
	return u.Deleted
}

func (u *UserModel) HasEmailVerified() bool {
	return u.EmailVerified
}
