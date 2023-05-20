package types

import (
	"errors"
	"net/mail"
	"strings"
)

type Email string

func ParseAndValidateEmail(email string) (Email, error) {
	email = strings.ToLower(email)

	parsedEmail, err := mail.ParseAddress(email)

	if err != nil {
		return "", errors.New("INVALID_EMAIL")
	}

	if len(parsedEmail.Address) > 100 {
		return "", errors.New("INVALID_EMAIL")
	}

	return Email(parsedEmail.Address), nil
}
