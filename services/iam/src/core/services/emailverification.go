package services

import (
	"errors"
	"fmt"
	"iam/pkg/apperrors"
	"iam/src/core/domain/model"
	"iam/src/core/domain/types"
	"iam/src/core/ports/primaryports"
	"iam/src/core/ports/secondaryports"
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

func (e *EmailVerificationService) Send(args primaryports.SendVerificationCodeArgs) error {
	userId, err := types.ParseAndValidateId(args.UserId)

	if err != nil {
		return err
	}

	user, err := e.usersRepository.GetUserById(userId)

	if err != nil {
		return err
	}

	if user.HasEmailVerified() {
		return errors.New(apperrors.EmailAlreadyVerified)
	}

	nActiveCodes, err := e.emailVerificationCodeRepository.CountActiveCodes(userId)

	if err != nil {
		return err
	}

	if nActiveCodes >= 3 {
		return errors.New(apperrors.LimitExceeded)
	}

	verificationCode, err := model.NewVerificationCodeModel(userId)

	if err != nil {
		return err
	}

	err = e.emailVerificationCodeRepository.SaveCode(verificationCode)

	if err != nil {
		return err
	}

	fmt.Println("==== Sending email ====", user.Email, verificationCode.Code, verificationCode.Code.EncodeForURL())

	return nil
}

func (e *EmailVerificationService) Confirm(args primaryports.ConfirmEmailArgs) (primaryports.ConfirmEmailAnswer, error) {
	code, err := types.ParseAndValidateCode(args.VerificationCode)

	if err != nil {
		return primaryports.ConfirmEmailAnswer{}, err
	}

	verificationCode, err := e.emailVerificationCodeRepository.GetCode(code)

	if err != nil {
		return primaryports.ConfirmEmailAnswer{}, err
	}

	if verificationCode.IsExpired() {
		return primaryports.ConfirmEmailAnswer{}, errors.New(apperrors.VerificationCodeExpired)
	}

	err = e.emailVerificationCodeRepository.DeleteCode(code)

	if err != nil {
		return primaryports.ConfirmEmailAnswer{}, err
	}

	user, err := e.usersRepository.GetUserById(verificationCode.UserId)

	if err != nil {
		return primaryports.ConfirmEmailAnswer{}, err
	}

	if user.HasEmailVerified() {
		return primaryports.ConfirmEmailAnswer{}, errors.New(apperrors.EmailAlreadyVerified)
	}

	err = e.usersRepository.ValidateEmail(verificationCode.UserId)

	if err != nil {
		return primaryports.ConfirmEmailAnswer{}, err
	}

	if user.IsBlocked() {
		return primaryports.ConfirmEmailAnswer{}, errors.New(apperrors.UserBlocked)
	}

	if user.IsDeleted() {
		return primaryports.ConfirmEmailAnswer{}, errors.New(apperrors.UserDeleted)
	}

	authorizationCode, err := model.NewAuthorizationCodeModel(verificationCode.UserId)

	if err != nil {
		return primaryports.ConfirmEmailAnswer{}, err
	}

	err = e.authorizationCodeRepository.SaveCode(authorizationCode)

	if err != nil {
		return primaryports.ConfirmEmailAnswer{}, err
	}

	return primaryports.ConfirmEmailAnswer{
		AuthorizationCode: string(authorizationCode.Code),
	}, nil
}
