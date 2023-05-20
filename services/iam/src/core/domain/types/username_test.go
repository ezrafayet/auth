package types

import (
	"testing"
)

func TestUsername_ParseAndValidateUsername_Valid(t *testing.T) {
	validUsernames := []string{
		"foobar",
		"foo",
		"foo-bar",
		"FooBar",
		"foo_bar",
	}

	for _, username := range validUsernames {
		_, err := ParseAndValidateUsername(username)

		if err != nil {
			t.Errorf("Expected username to be valid, got error %v for username: %v", err, username)
		}
	}
}

func TestUsername_ParseAndValidateUsername_Invalid(t *testing.T) {
	invalidUsernames := []string{
		"",
		"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
		"foobar!",
		"foobar#",
		"foobar$",
		"-foobar",
		"_foobar",
		"foobar-",
		"foobar_",
		"foo bar",
		" foobar",
	}

	for _, username := range invalidUsernames {
		_, err := ParseAndValidateUsername(username)

		if err == nil {
			t.Errorf("Expected an error for username: %v", username)
		}
	}
}

func TestUsername_ComputeUsernameFingerprint_Expected(t *testing.T) {
	username := Username("aa")

	fingerprint := ComputeUsernameFingerprint(username)

	fingerprintExpected := "961b6dd3ede3cb8ecbaacbd68de040cd78eb2ed5889130cceb4c49268ea4d506"

	if string(fingerprint) != fingerprintExpected {
		t.Errorf("Expected %v, got %v", fingerprintExpected, fingerprint)
	}
}

func TestUsername_ComputeUsernameFingerprint_Consistency(t *testing.T) {
	username := Username("aa")

	fingerprint1 := ComputeUsernameFingerprint(username)

	fingerprint2 := ComputeUsernameFingerprint(username)

	equal := fingerprint1 == fingerprint2

	if !equal {
		t.Errorf("Expected true, got %v", equal)
	}
}
