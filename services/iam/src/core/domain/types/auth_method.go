package types

import (
	"errors"
	"iam/pkg/apperrors"
)

type AuthMethod int8

const (
	AuthMethodPassword AuthMethod = iota
	AuthMethodMagicLink
)

var authMethodsNames = []string{"password", "magic-link"}

func (a AuthMethod) String() string {
	if a < AuthMethodPassword || a > AuthMethodMagicLink {
		return "unknown"
	}
	return authMethodsNames[a]
}

func ParseAndValidateAuthMethod(authMethod string) (AuthMethod, error) {
	switch authMethod {
	case authMethodsNames[AuthMethodPassword]:
		return AuthMethodPassword, nil
	case authMethodsNames[AuthMethodMagicLink]:
		return AuthMethodMagicLink, nil
	default:
		return 127, errors.New(apperrors.InvalidAuthMethod)
	}
}
