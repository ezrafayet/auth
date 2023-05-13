package types

import "github.com/google/uuid"

type Id string

func NewId() Id {
	return Id(uuid.New().String())
}

func IsIdValid(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
