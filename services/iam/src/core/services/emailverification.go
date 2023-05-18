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
	authorizationCodeRepository     secondaryports.AuthorizationCodeRepository
}

func NewEmailVerificationService(
	usersRepository secondaryports.UsersRepository,
	emailVerificationCodeRepository secondaryports.EmailVerificationCodeRepository,
	authorizationCodeRepository secondaryports.AuthorizationCodeRepository) *EmailVerificationService {
	return &EmailVerificationService{
		usersRepository:                 usersRepository,
		emailVerificationCodeRepository: emailVerificationCodeRepository,
		authorizationCodeRepository:     authorizationCodeRepository,
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

	if user.HasEmailVerified() {
		return errors.New("EMAIL_ALREADY_VERIFIED")
	}

	nActiveCodes, err := e.emailVerificationCodeRepository.CountActiveCodes(userId)

	if err != nil {
		fmt.Println(err)
		return errors.New("SERVER_ERROR")
	}

	if nActiveCodes >= 3 {
		return errors.New("LIMIT_EXCEEDED")
	}

	verificationCode, err := model.NewVerificationCodeModel(userId)

	if err != nil {
		return err
	}

	err = e.emailVerificationCodeRepository.SaveCode(verificationCode)

	if err != nil {
		fmt.Println(err)
		return errors.New("SERVER_ERROR")
	}

	fmt.Println("==== Sending email ====", user.Email, verificationCode.Code, verificationCode.Code.GetUrlEncoded())

	return nil
}

func (e *EmailVerificationService) Confirm(args ConfirmEmailArgs) (ConfirmEmailAnswer, error) {
	code, err := types.ParseAndValidateCode(args.VerificationCode)

	if err != nil {
		return ConfirmEmailAnswer{}, err
	}

	verificationCode, err := e.emailVerificationCodeRepository.GetCode(code)

	if err != nil {
		return ConfirmEmailAnswer{}, err
	}

	if verificationCode.IsExpired() {
		return ConfirmEmailAnswer{}, errors.New("CODE_EXPIRED")
	}

	err = e.emailVerificationCodeRepository.DeleteCode(code)

	if err != nil {
		return ConfirmEmailAnswer{}, err
	}

	user, err := e.usersRepository.GetUserById(verificationCode.UserId)

	if user.HasEmailVerified() {
		return ConfirmEmailAnswer{}, errors.New("EMAIL_ALREADY_VERIFIED")
	}

	err = e.usersRepository.ValidateEmail(verificationCode.UserId)

	if err != nil {
		return ConfirmEmailAnswer{}, err
	}

	authorizationCode, err := model.NewAuthorizationCodeModel(verificationCode.UserId)

	if err != nil {
		return ConfirmEmailAnswer{}, err
	}

	err = e.authorizationCodeRepository.SaveCode(authorizationCode)

	if err != nil {
		return ConfirmEmailAnswer{}, err
	}

	return ConfirmEmailAnswer{
		AuthorizationCode: string(authorizationCode.Code),
	}, nil
}
