package services

import (
	"errors"
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
	emailRepository                 secondaryports.EmailRepository
}

func NewEmailVerificationService(
	usersRepository secondaryports.UsersRepository,
	emailVerificationCodeRepository secondaryports.EmailVerificationCodeRepository,
	authorizationCodeRepository secondaryports.AuthorizationCodeRepository,
	emailRepository secondaryports.EmailRepository) *EmailVerificationService {
	return &EmailVerificationService{
		usersRepository:                 usersRepository,
		emailVerificationCodeRepository: emailVerificationCodeRepository,
		authorizationCodeRepository:     authorizationCodeRepository,
		emailRepository:                 emailRepository,
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

	verificationCode, err := model.NewEmailVerificationCode(userId)

	if err != nil {
		return err
	}

	err = e.emailVerificationCodeRepository.SaveCode(verificationCode)

	if err != nil {
		return err
	}

	go func() {
		_ = e.emailRepository.SendVerificationCode(user.Email, user.Username, verificationCode.Code.EncodeForURL())
	}()

	return nil
}

func (e *EmailVerificationService) Confirm(args primaryports.ConfirmEmailArgs) (primaryports.ConfirmEmailAnswer, error) {
	code, err := types.ParseAndValidateCode(args.VerificationCode)

	if err != nil {
		return primaryports.ConfirmEmailAnswer{}, err
	}

	verificationCode, err := e.emailVerificationCodeRepository.GetCode(code) // turn this into a get and delete ?

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

	authorizationCode, err := model.NewAuthorizationCode(verificationCode.UserId)

	if err != nil {
		return primaryports.ConfirmEmailAnswer{}, err
	}

	err = e.authorizationCodeRepository.SaveCode(authorizationCode)

	if err != nil {
		return primaryports.ConfirmEmailAnswer{}, err
	}

	return primaryports.ConfirmEmailAnswer{
		// here the code does not need to be encoded because it is not used in a URL
		AuthorizationCode: string(authorizationCode.Code),
	}, nil
}
