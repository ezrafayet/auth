package types

import (
	"errors"
	"iam/pkg/apperrors"
)

type AuthType int8

const (
	AuthTypePassword AuthType = iota + 1
	AuthTypeMagicLink
)

var authTypesNames = map[AuthType]string{
	AuthTypePassword:  "password",
	AuthTypeMagicLink: "magic-link",
}

func (a AuthType) String() string {
	if a < AuthTypePassword || a > AuthTypeMagicLink {
		return "unknown"
	}

	return authTypesNames[a]
}

func ParseAndValidateAuthType(authMethod string) (AuthType, error) {
	switch authMethod {
	case authTypesNames[AuthTypePassword]:
		return AuthTypePassword, nil
	case authTypesNames[AuthTypeMagicLink]:
		return AuthTypeMagicLink, nil
	default:
		return 0, errors.New(apperrors.InvalidAuthType)
	}
}
