package services

import (
	"iam/src/core/domain/model"
	"iam/src/core/domain/types"
	"iam/src/core/ports/primaryports"
	"iam/src/core/ports/secondaryports"
)

type UsersService struct {
	usersRepository secondaryports.UsersRepository
}

func NewUserService(usersRepository secondaryports.UsersRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}

func (s *UsersService) Register(args primaryports.RegisterArgs) (primaryports.RegisterAnswer, error) {
	authMethod, err := types.ParseAndValidateAuthMethod(args.AuthMethod)

	if err != nil {
		return primaryports.RegisterAnswer{}, err
	}

	email, err := types.ParseAndValidateEmail(args.Email)

	if err != nil {
		return primaryports.RegisterAnswer{}, err
	}

	username, err := types.ParseAndValidateUsername(args.Username)

	if err != nil {
		return primaryports.RegisterAnswer{}, err
	}

	user := model.NewUserModel(username, email, types.NewTimestamp())

	if err != nil {
		return primaryports.RegisterAnswer{}, err
	}

	userAuthMethod := model.NewUsersAuthMethodsModel(user.Id, authMethod)

	err = s.usersRepository.SaveUser(user, userAuthMethod)

	if err != nil {
		return primaryports.RegisterAnswer{}, err
	}

	return primaryports.RegisterAnswer{
		UserId: user.Id,
	}, nil
}
