package types

import (
	"crypto/sha256"
	"errors"
	"strings"
)

type Username string
type UsernameFingerprint string

func ParseAndValidateUsername(username string) (Username, error) {
	username = strings.TrimSpace(username)
	if len(username) > 40 || len(username) < 2 {
		return "", errors.New("INVALID_USERNAME")
	}
	if username[0] == '-' || username[len(username)-1] == '-' {
		return "", errors.New("NO_TRAILING_OR_LEADING_DASH")
	}
	authorizedCharacters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	for _, char := range username {
		if !strings.ContainsRune(authorizedCharacters, char) {
			return "", errors.New("INVALID_USERNAME")
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
	return UsernameFingerprint(sha[:])
}
