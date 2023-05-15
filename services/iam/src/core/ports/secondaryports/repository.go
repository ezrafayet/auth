package secondaryports

import (
	"iam/src/core/model"
	"iam/src/core/types"
)

type UsersRepository interface {
	SaveUser(user model.UserModel, authMethod model.UsersAuthMethodsModel) error
	GetUserById(id types.Id) (model.UserModel, error)
	ValidateEmail(userId types.Id) error
}

type EmailVerificationCodeRepository interface {
	SaveVerificationCode(code model.EmailVerificationCodeModel) error
	ConfirmVerificationCode(code types.Code) error
}
