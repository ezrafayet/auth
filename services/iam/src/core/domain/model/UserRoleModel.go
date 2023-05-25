package model

import (
	"iam/src/core/domain/types"
)

type UserRoleModel struct {
	UserId    types.Id
	RoleId    int8
	CreatedAt types.Timestamp
}

func NewUserRoleModel(userId types.Id, role types.Role) UserRoleModel {
	return UserRoleModel{
		UserId:    userId,
		RoleId:    int8(role),
		CreatedAt: types.NewTimestamp(),
	}
}
