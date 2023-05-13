package ports

import (
	"iam/src/core/model"
	"iam/src/core/types"
)

type UsersRepository interface {
	CreateUser(user model.UserModel, authMethod types.AuthMethod) error
	//GetUserById(id types.Id) (model.UserModel, error)
}

type VerificationCodeRepository interface {
}
