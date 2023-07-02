package services

import (
	"errors"
	"time"

	"iam/pkg/apperrors"
	"iam/src/core/domain/model"
	"iam/src/core/domain/types"
	"iam/src/core/ports/primaryports"
	"iam/src/core/ports/secondaryports"
)

type AuthenticationService struct {
	usersRepository             secondaryports.UsersRepository
	authorizationCodeRepository secondaryports.AuthorizationCodeRepository
	refreshTokenRepository      secondaryports.RefreshTokenRepository
	emailRepository             secondaryports.EmailRepository
}

func NewAuthenticationService(
	usersRepository secondaryports.UsersRepository,
	authorizationCodeRepository secondaryports.AuthorizationCodeRepository,
	emailRepository secondaryports.EmailRepository,
	refreshTokenRepository secondaryports.RefreshTokenRepository) *AuthenticationService {
	return &AuthenticationService{
		usersRepository:             usersRepository,
		authorizationCodeRepository: authorizationCodeRepository,
		emailRepository:             emailRepository,
		refreshTokenRepository:      refreshTokenRepository,
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

	// todo: count existing authorization codes and block if too many

	authorizationCode, err := model.NewAuthorizationCode(user.Id)

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

func (a *AuthenticationService) Authenticate(args primaryports.GetAccessTokenArgs) (primaryports.GetAccessTokenAnswer, error) {
	code, err := types.ParseAndValidateCode(args.AuthorizationCode)

	if err != nil {
		return primaryports.GetAccessTokenAnswer{}, err
	}

	authorizationCode, err := a.authorizationCodeRepository.GetCode(code)

	if err != nil {
		return primaryports.GetAccessTokenAnswer{}, err
	}

	if authorizationCode.IsExpired() {
		return primaryports.GetAccessTokenAnswer{}, errors.New(apperrors.AuthorizationCodeExpired)
	}

	user, err := a.usersRepository.GetUserById(authorizationCode.UserId)

	if err != nil {
		return primaryports.GetAccessTokenAnswer{}, err
	}

	if !user.HasEmailVerified() {
		return primaryports.GetAccessTokenAnswer{}, errors.New(apperrors.EmailNotVerified)
	}

	if user.IsBlocked() {
		return primaryports.GetAccessTokenAnswer{}, errors.New(apperrors.UserBlocked)
	}

	if user.IsDeleted() {
		return primaryports.GetAccessTokenAnswer{}, errors.New(apperrors.UserDeleted)
	}

	accessToken, expiresAt, err := types.NewAccessToken(types.CustomClaims{
		UserId: string(user.Id),
		Roles:  "user",
	}, time.Now().UTC())

	if err != nil {
		return primaryports.GetAccessTokenAnswer{}, err
	}

	refreshToken, err := model.NewRefreshToken(user.Id)

	if err != nil {
		return primaryports.GetAccessTokenAnswer{}, err
	}

	err = a.authorizationCodeRepository.DeleteCode(authorizationCode.Code)

	if err != nil {
		return primaryports.GetAccessTokenAnswer{}, err
	}

	err = a.refreshTokenRepository.SaveToken(refreshToken)

	if err != nil {
		return primaryports.GetAccessTokenAnswer{}, err
	}

	return primaryports.GetAccessTokenAnswer{
		AccessToken:  string(accessToken),
		RefreshToken: string(refreshToken.Token),
		ExpiresAt:    expiresAt,
	}, nil
}
