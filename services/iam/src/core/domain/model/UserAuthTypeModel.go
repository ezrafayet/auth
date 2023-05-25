package model

import "iam/src/core/domain/types"

type UserAuthTypeModel struct {
	UserId     types.Id
	AuthTypeId int8
	CreatedAt  types.Timestamp
}

func NewUserAuthTypeModel(
	userId types.Id, authMethodId types.AuthType) UserAuthTypeModel {
	return UserAuthTypeModel{
		UserId:     userId,
		AuthTypeId: int8(authMethodId),
		CreatedAt:  types.NewTimestamp(),
	}
}
