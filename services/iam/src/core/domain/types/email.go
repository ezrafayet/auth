package types

import (
	"errors"
	"iam/pkg/apperrors"
	"net/mail"
	"strings"
)

type Email string

func ParseAndValidateEmail(email string) (Email, error) {
	email = strings.ToLower(email)

	parsedEmail, err := mail.ParseAddress(email)

	if err != nil {
		return "", err
	}

	if len(parsedEmail.Address) > 100 {
		return "", errors.New(apperrors.InvalidEmail)
	}

	return Email(parsedEmail.Address), nil
}
