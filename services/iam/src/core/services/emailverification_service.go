package services

import (
	"iam/src/core/ports/secondaryports"
)

type EmailVerificationService struct {
	usersRepository            secondaryports.UsersRepository
	verificationCodeRepository secondaryports.VerificationCodeRepository
}
