package types

import (
	"errors"
	"iam/pkg/apperrors"
)

type Role int8

const (
	RoleAdmin Role = iota
	RoleUser
)

var rolesNames = []string{"admin", "user"}

func (r Role) String() string {
	if r < RoleAdmin || r > RoleUser {
		return "unknown"
	}

	return rolesNames[r]
}

func ParseAndValidateRole(role string) (Role, error) {
	switch role {
	case rolesNames[RoleAdmin]:
		return RoleAdmin, nil
	case rolesNames[RoleUser]:
		return RoleUser, nil
	default:
		return 127, errors.New(apperrors.InvalidRole)
	}
}
