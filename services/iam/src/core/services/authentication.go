package services

import (
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
	return primaryports.SendMagicLinkAnswer{}, nil
}
