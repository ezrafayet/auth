package services

import "iam/src/core/domain/types"

type RegisterArgs struct {
	AuthMethod string `json:"authMethod"`
	Email      string `json:"email"`
	Username   string `json:"username"`
}

type RegisterAnswer struct {
	UserId types.Id `json:"userId"`
}
