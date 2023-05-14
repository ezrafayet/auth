package types

import (
	"crypto/rand"
	"encoding/base64"
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
