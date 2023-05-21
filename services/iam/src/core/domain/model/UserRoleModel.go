package model

import (
	"iam/src/core/domain/types"
)

type UserRoleModel struct {
	Id     types.Id
	UserId types.Id
	RoleId types.Role
}

func NewUserRoleModel(userId types.Id, role types.Role) UserRoleModel {
	return UserRoleModel{
		Id:     types.NewId(),
		UserId: userId,
		RoleId: role,
	}
}
