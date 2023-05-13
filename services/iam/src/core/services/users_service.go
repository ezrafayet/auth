package services

import (
	"iam/src/core/model"
	"iam/src/core/ports/secondaryports"
	"iam/src/core/types"
)

type UsersService struct {
	usersRepository secondaryports.UsersRepository
}

func NewUserService(usersRepository secondaryports.UsersRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}

func (s *UsersService) Register(args RegisterArgs) (RegisterAnswer, error) {
	authMethod, err := types.ParseAndValidateAuthMethod(args.AuthMethod)

	if err != nil {
		return RegisterAnswer{}, err
	}

	email, err := types.ParseAndValidateEmail(args.Email)

	if err != nil {
		return RegisterAnswer{}, err
	}

	username, err := types.ParseAndValidateUsername(args.Username)

	if err != nil {
		return RegisterAnswer{}, err
	}

	user := model.NewUserModel(username, email)

	if err != nil {
		return RegisterAnswer{}, err
	}

	err = s.usersRepository.CreateUser(user, authMethod)

	if err != nil {
		return RegisterAnswer{}, err
	}

	return RegisterAnswer{
		UserId: user.Id,
	}, nil
}
