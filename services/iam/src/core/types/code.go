package types

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

// Code is a generic code type, used for verification codes, authorization codes, etc.
// It is a base64 encoded string of 32 random bytes. It is not a hash, and it has 44 characters.
type Code string

func NewCode() (Code, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return Code(base64.StdEncoding.EncodeToString(b)), nil
}

func ParseAndValidateCode(code string) (Code, error) {
	if len(code) != 44 {
		return "", errors.New("INVALID_CODE")
	}

	_, err := base64.StdEncoding.DecodeString(code)

	if err != nil {
		return "", errors.New("INVALID_CODE")
	}

	return Code(code), nil
}

func (c Code) GetUrlEncoded() string {
	return base64.URLEncoding.EncodeToString([]byte(c))
}
