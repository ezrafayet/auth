package model

import "iam/src/core/domain/types"

type UsersAuthMethodsModel struct {
	Id         types.Id
	UserId     types.Id
	AuthMethod types.AuthMethod
}

func NewUsersAuthMethodsModel(
	userId types.Id, authMethodId types.AuthMethod) UsersAuthMethodsModel {
	return UsersAuthMethodsModel{
		Id:         types.NewId(),
		UserId:     userId,
		AuthMethod: authMethodId,
	}
}
