package services

import (
	"errors"
	"iam/pkg/apperrors"
	"iam/src/core/domain/model"
	"iam/src/core/domain/types"
	"iam/src/core/ports/primaryports"
	"iam/src/core/ports/secondaryports"
)

type UsersService struct {
	usersRepository secondaryports.UsersRepository
	emailRepository secondaryports.EmailRepository
}

// NewUserService initializes a new instance of UsersService with the provided repositories
func NewUserService(usersRepository secondaryports.UsersRepository, emailRepository secondaryports.EmailRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
		emailRepository: emailRepository,
	}
}

// Register registers a new user
func (s *UsersService) Register(args primaryports.RegisterArgs) (primaryports.RegisterAnswer, error) {
	authMethod, err := types.ParseAndValidateAuthMethod(args.AuthMethod)

	if err != nil {
		return primaryports.RegisterAnswer{}, err
	}

	email, err := types.ParseAndValidateEmail(args.Email)

	if err != nil {
		switch err.Error() {
		case apperrors.InvalidEmail:
			return primaryports.RegisterAnswer{}, err
		default:
			return primaryports.RegisterAnswer{}, errors.New(apperrors.InvalidEmail)
		}
	}

	username, err := types.ParseAndValidateUsername(args.Username)

	if err != nil {
		return primaryports.RegisterAnswer{}, err
	}

	user := model.NewUserModel(username, email, types.NewTimestamp())

	userAuthMethod := model.NewUsersAuthMethodsModel(user.Id, authMethod)

	err = s.usersRepository.SaveUser(user, userAuthMethod)

	if err != nil {
		return primaryports.RegisterAnswer{}, err
	}

	go func() {
		_ = s.emailRepository.WelcomeNewUser(user.Email, user.Username)
	}()

	return primaryports.RegisterAnswer{
		UserId: user.Id,
	}, nil
}
