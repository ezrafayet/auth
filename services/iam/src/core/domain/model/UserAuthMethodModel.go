package model

import "iam/src/core/domain/types"

type UserAuthMethodModel struct {
	Id         types.Id
	UserId     types.Id
	AuthMethod types.AuthMethod
}

func NewUsersAuthMethodsModel(
	userId types.Id, authMethodId types.AuthMethod) UserAuthMethodModel {
	return UserAuthMethodModel{
		Id:         types.NewId(),
		UserId:     userId,
		AuthMethod: authMethodId,
	}
}
