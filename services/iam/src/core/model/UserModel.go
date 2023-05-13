package model

import (
	"iam/src/core/types"
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

func NewUserModel(username types.Username, email types.Email) UserModel {
	return UserModel{
		Id:                  types.NewId(),
		CreatedAt:           types.NewTimestamp(),
		Username:            username,
		UsernameFingerprint: types.ComputeUsernameFingerprint(username),
		Email:               email,
		EmailVerified:       false,
		Blocked:             false,
		Deleted:             false,
	}
}

func (u *UserModel) Hydrate() { /*todo*/ }

func (u *UserModel) IsBlocked() bool {
	return u.Blocked
}

func (u *UserModel) IsDeleted() bool {
	return u.Deleted
}

func (u *UserModel) IsEmailVerified() bool {
	return u.EmailVerified
}
