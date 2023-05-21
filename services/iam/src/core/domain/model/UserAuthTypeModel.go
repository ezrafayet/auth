package model

import "iam/src/core/domain/types"

type UserAuthTypeModel struct {
	Id       types.Id
	UserId   types.Id
	AuthType types.AuthType
}

func NewUserAuthTypeModel(
	userId types.Id, authMethodId types.AuthType) UserAuthTypeModel {
	return UserAuthTypeModel{
		Id:       types.NewId(),
		UserId:   userId,
		AuthType: authMethodId,
	}
}
