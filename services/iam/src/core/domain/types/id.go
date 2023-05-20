package types

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"iam/pkg/apperrors"
)

type Id string

func NewId() Id {
	return Id(uuid.New().String())
}

func ParseAndValidateId(id string) (Id, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		fmt.Println(err)
		return "", errors.New(apperrors.InvalidId)
	}
	return Id(id), nil
}
