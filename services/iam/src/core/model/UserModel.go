package model

import (
	"iam/src/core/types"
	"time"
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

func NewUserModel(username types.Username, email types.Email, timestamp types.Timestamp) UserModel {
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

func (u *UserModel) Hydrate(
	id string,
	createdAt time.Time,
	username string,
	usernameFingerprint string,
	email string,
	emailVerified bool,
	emailVerifiedAt time.Time,
	blocked bool,
	deleted bool,
	deletedAt time.Time) {
	u.Id = types.Id(id)
	u.CreatedAt = types.Timestamp(createdAt)
	u.Username = types.Username(username)
	u.UsernameFingerprint = types.UsernameFingerprint(usernameFingerprint)
	u.Email = types.Email(email)
	u.EmailVerified = emailVerified
	u.EmailVerifiedAt = types.Timestamp(emailVerifiedAt)
	u.Blocked = blocked
	u.Deleted = deleted
	u.DeletedAt = types.Timestamp(deletedAt)
}

func (u *UserModel) IsBlocked() bool {
	return u.Blocked
}

func (u *UserModel) IsDeleted() bool {
	return u.Deleted
}

func (u *UserModel) IsEmailVerified() bool {
	return u.EmailVerified
}
