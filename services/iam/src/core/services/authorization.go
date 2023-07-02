package services

import (
	"errors"
	"iam/pkg/apperrors"
	"iam/src/core/domain/model"
	"iam/src/core/domain/types"
	"iam/src/core/ports/primaryports"
	"iam/src/core/ports/secondaryports"
	"os"
)

type AuthorizationService struct {
	userRepository              secondaryports.UsersRepository
	authorizationCodeRepository secondaryports.AuthorizationCodeRepository
	refreshTokenRepository      secondaryports.RefreshTokenRepository
}

func NewAuthorizationService(
	userRepository secondaryports.UsersRepository,
	authorizationCodeRepository secondaryports.AuthorizationCodeRepository,
	refreshTokenRepository secondaryports.RefreshTokenRepository) *AuthorizationService {
	return &AuthorizationService{
		userRepository:              userRepository,
		authorizationCodeRepository: authorizationCodeRepository,
		refreshTokenRepository:      refreshTokenRepository,
	}
}

func (a *AuthorizationService) AreFeaturesEnabled(args primaryports.AreFeaturesEnabledArgs) (primaryports.AreFeaturesEnabledAnswer, error) {
	for _, f := range args.FlagsNeeded {
		if os.Getenv(f) != "true" {
			return primaryports.AreFeaturesEnabledAnswer{Active: false}, nil
		}
	}

	return primaryports.AreFeaturesEnabledAnswer{Active: true}, nil
}

func (a *AuthorizationService) IsAccessTokenValid(args primaryports.IsAccessTokenValidArgs) (primaryports.IsAccessTokenValidAnswer, error) {
	if args.AuthorisationHeader == "" {
		return primaryports.IsAccessTokenValidAnswer{Valid: false}, nil
	}

	if args.AuthorisationHeader[:7] != "Bearer " {
		return primaryports.IsAccessTokenValidAnswer{Valid: false}, nil
	}

	extractedToken := args.AuthorisationHeader[7:]

	ok, _, err := types.ParseAndValidateAccessToken(extractedToken)

	if err != nil {
		return primaryports.IsAccessTokenValidAnswer{Valid: false}, err
	}

	return primaryports.IsAccessTokenValidAnswer{
		Valid: ok,
	}, nil
}

func (a *AuthorizationService) RefreshAccessToken(args primaryports.RefreshAccessTokenArgs) (primaryports.RefreshAccessTokenAnswer, error) {
	strRefreshToken, err := types.ParseAndValidateCode(args.RefreshToken)

	if err != nil {
		return primaryports.RefreshAccessTokenAnswer{}, err
	}

	refreshToken, err := a.refreshTokenRepository.GetAndDeleteByToken(strRefreshToken)

	if err != nil {
		return primaryports.RefreshAccessTokenAnswer{}, err
	}

	if refreshToken.IsExpired() {
		return primaryports.RefreshAccessTokenAnswer{}, errors.New(apperrors.InvalidRefreshToken)
	}

	user, err := a.userRepository.GetUserById(refreshToken.UserId)

	if err != nil {
		return primaryports.RefreshAccessTokenAnswer{}, err
	}

	if !user.HasEmailVerified() {
		return primaryports.RefreshAccessTokenAnswer{}, errors.New(apperrors.EmailNotVerified)
	}

	if user.IsBlocked() {
		return primaryports.RefreshAccessTokenAnswer{}, errors.New(apperrors.UserBlocked)
	}

	if user.IsDeleted() {
		return primaryports.RefreshAccessTokenAnswer{}, errors.New(apperrors.UserDeleted)
	}

	// count existing tokens

	authorizationCode, err := model.NewAuthorizationCode(user.Id)

	if err != nil {
		return primaryports.RefreshAccessTokenAnswer{}, err
	}

	err = a.authorizationCodeRepository.SaveCode(authorizationCode)

	if err != nil {
		return primaryports.RefreshAccessTokenAnswer{}, err
	}

	return primaryports.RefreshAccessTokenAnswer{
		AuthorizationCode: string(authorizationCode.Code),
	}, nil
}

func (a *AuthorizationService) IsCaptchaValid(args primaryports.IsCaptchaValidArgs) (primaryports.IsCaptchaValidAnswer, error) {
	if os.Getenv("ENABLE_CAPTCHA") != "true" {
		return primaryports.IsCaptchaValidAnswer{Valid: true}, nil
	}

	if args.CaptchaResponse == "" {
		return primaryports.IsCaptchaValidAnswer{Valid: false}, nil
	}

	// todo: implement captcha validation

	return primaryports.IsCaptchaValidAnswer{Valid: false}, nil
}
