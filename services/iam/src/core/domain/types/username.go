package types

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"iam/pkg/apperrors"
	"strings"
)

type Username string
type UsernameFingerprint string

func ParseAndValidateUsername(username string) (Username, error) {
	if len(username) > 40 || len(username) < 2 {
		return "", errors.New(apperrors.InvalidUsername)
	}
	if username[0] == '-' || username[len(username)-1] == '-' {
		return "", errors.New(apperrors.InvalidUsername)
	}
	if username[0] == '_' || username[len(username)-1] == '_' {
		return "", errors.New(apperrors.InvalidUsername)
	}
	authorizedCharacters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	for _, char := range username {
		if !strings.ContainsRune(authorizedCharacters, char) {
			return "", errors.New(apperrors.InvalidUsername)
		}
	}
	return Username(username), nil
}

func ComputeUsernameFingerprint(username Username) UsernameFingerprint {
	usernameStr := strings.TrimSpace(string(username))
	usernameStr = strings.ToLower(usernameStr)
	var cleanedUsername string
	authorizedLowerCharacters := "abcdefghijklmnopqrstuvwxyz0123456789"
	for _, cha := range usernameStr {
		if strings.ContainsRune(authorizedLowerCharacters, cha) {
			cleanedUsername += string(cha)
		} else {
			cleanedUsername += "$"
		}
	}
	sha := sha256.Sum256([]byte(cleanedUsername))
	return UsernameFingerprint(fmt.Sprintf("%x", sha[:]))
}
