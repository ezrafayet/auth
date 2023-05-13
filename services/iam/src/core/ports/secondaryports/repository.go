package secondaryports

import (
	"iam/src/core/model"
)

type UsersRepository interface {
	CreateUser(user model.UserModel, authMethod model.UsersAuthMethodsModel) error
	//GetUserById(id types.Id) (model.UserModel, error)
}

type VerificationCodeRepository interface {
}
