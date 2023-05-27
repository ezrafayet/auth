package services

import (
	"iam/src/core/domain/types"
	"iam/src/core/ports/primaryports"
)

type AuthorizationService struct{}

func NewAuthorizationService() *AuthorizationService {
	return &AuthorizationService{}
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
