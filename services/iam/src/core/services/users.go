package services

import (
	"fmt"
	"iam/src/core/model"
	"iam/src/core/ports"
	"iam/src/core/types"
)

type UsersService struct {
	usersRepository ports.UsersRepository
}

func NewUserService(usersRepository ports.UsersRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}

func (s *UsersService) Register(args ports.RegisterArgs) (ports.RegisterAnswer, error) {
	authMethod, err := types.ParseAndValidateAuthMethod(args.AuthMethod)

	fmt.Println("===========", args)

	if err != nil {
		return ports.RegisterAnswer{}, err
	}

	email, err := types.ParseAndValidateEmail(args.Email)

	if err != nil {
		return ports.RegisterAnswer{}, err
	}

	username, err := types.ParseAndValidateUsername(args.Username)

	if err != nil {
		return ports.RegisterAnswer{}, err
	}

	user := model.NewUserModel(username, email)

	if err != nil {
		return ports.RegisterAnswer{}, err
	}

	err = s.usersRepository.CreateUser(user, authMethod)

	if err != nil {
		return ports.RegisterAnswer{}, err
	}

	return ports.RegisterAnswer{
		UserId: user.Id,
	}, nil
}
