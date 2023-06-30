package types

import (
	"errors"
	"iam/pkg/apperrors"
)

type Role int8

const (
	RoleAdmin Role = iota + 1
	RoleUser
)

// var rolesNames = []string{"admin", "user"}

var rolesNames = map[Role]string{
	RoleAdmin: "admin",
	RoleUser:  "user",
}

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
		return 0, errors.New(apperrors.InvalidRole)
	}
}
