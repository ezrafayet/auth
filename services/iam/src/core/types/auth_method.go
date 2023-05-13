package types

import "errors"

type AuthMethod int8

const (
	AuthMethodPassword AuthMethod = iota
	AuthMethodMagicLink
)

var AuthMethodsNames = []string{"password", "magic-link"}

func (a AuthMethod) String() string {
	return AuthMethodsNames[a]
}

func ParseAndValidateAuthMethod(authMethod string) (AuthMethod, error) {
	switch authMethod {
	case AuthMethodsNames[0]:
		return 0, nil
	case AuthMethodsNames[1]:
		return 1, nil
	default:
		return -1, errors.New("INVALID_AUTH_METHOD")
	}
}
