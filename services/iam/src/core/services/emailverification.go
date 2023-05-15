package services

import (
	"errors"
	"fmt"
	"iam/src/core/model"
	"iam/src/core/ports/secondaryports"
	"iam/src/core/types"
)

type EmailVerificationService struct {
	usersRepository                 secondaryports.UsersRepository
	emailVerificationCodeRepository secondaryports.EmailVerificationCodeRepository
}

func NewEmailVerificationService(
	usersRepository secondaryports.UsersRepository,
	emailVerificationCodeRepository secondaryports.EmailVerificationCodeRepository) *EmailVerificationService {
	return &EmailVerificationService{
		usersRepository:                 usersRepository,
		emailVerificationCodeRepository: emailVerificationCodeRepository,
	}
}

func (e *EmailVerificationService) Send(args SendVerificationCodeArgs) error {
	userId, err := types.ParseAndValidateId(args.UserId)

	if err != nil {
		return err
	}

	user, err := e.usersRepository.GetUserById(userId)

	if err != nil {
		return err
	}

	if user.IsEmailVerified() {
		return errors.New("EMAIL_ALREADY_VERIFIED")
	}

	// todo: verify number of emails sent in the last 1 hour

	code, err := types.NewCode()

	if err != nil {
		fmt.Println(err)
		return errors.New("SERVER_ERROR")
	}

	verificationCode := model.NewVerificationCodeModel(userId, code, types.NewTimestamp(), 5*60)

	err = e.emailVerificationCodeRepository.SaveVerificationCode(verificationCode)

	if err != nil {
		fmt.Println(err)
		return errors.New("SERVER_ERROR")
	}

	fmt.Println("==== Sending email ====", user.Email, code)

	return nil
}

func (e *EmailVerificationService) Confirm(args ConfirmVerificationCodeArgs) (ConfirmVerificationCodeAnswer, error) {
	return ConfirmVerificationCodeAnswer{}, nil
}
