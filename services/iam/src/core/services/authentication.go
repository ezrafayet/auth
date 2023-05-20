package services

import (
	"errors"
	"iam/pkg/apperrors"
	"iam/src/core/domain/model"
	"iam/src/core/domain/types"
	"iam/src/core/ports/primaryports"
	"iam/src/core/ports/secondaryports"
)

type AuthenticationService struct {
	usersRepository             secondaryports.UsersRepository
	authorizationCodeRepository secondaryports.AuthorizationCodeRepository
	emailRepository             secondaryports.EmailRepository
}

func NewAuthenticationService(
	usersRepository secondaryports.UsersRepository,
	authorizationCodeRepository secondaryports.AuthorizationCodeRepository,
	emailRepository secondaryports.EmailRepository) *AuthenticationService {
	return &AuthenticationService{
		usersRepository:             usersRepository,
		authorizationCodeRepository: authorizationCodeRepository,
		emailRepository:             emailRepository,
	}
}

func (a *AuthenticationService) SendMagicLink(args primaryports.SendMagicLinkArgs) (primaryports.SendMagicLinkAnswer, error) {
	email, err := types.ParseAndValidateEmail(args.Email)

	if err != nil {
		return primaryports.SendMagicLinkAnswer{}, err
	}

	user, err := a.usersRepository.GetUserByEmail(email)

	if err != nil {
		return primaryports.SendMagicLinkAnswer{}, err
	}

	// todo: verify if user has magic-link enabled...

	if !user.HasEmailVerified() {
		return primaryports.SendMagicLinkAnswer{}, errors.New(apperrors.EmailNotVerified)
	}

	if user.IsBlocked() {
		return primaryports.SendMagicLinkAnswer{}, errors.New(apperrors.UserBlocked)
	}

	if user.IsDeleted() {
		return primaryports.SendMagicLinkAnswer{}, errors.New(apperrors.UserDeleted)
	}

	// todo: count existing authorization codes and delete them if there are too many
	// or delete all existing

	authorizationCode, err := model.NewAuthorizationCodeModel(user.Id)

	if err != nil {
		return primaryports.SendMagicLinkAnswer{}, err
	}

	err = a.authorizationCodeRepository.SaveCode(authorizationCode)

	if err != nil {
		return primaryports.SendMagicLinkAnswer{}, err
	}

	go func() {
		_ = a.emailRepository.SendMagicLink(user.Email, user.Username, authorizationCode.Code.EncodeForURL())
	}()

	return primaryports.SendMagicLinkAnswer{}, nil
}
