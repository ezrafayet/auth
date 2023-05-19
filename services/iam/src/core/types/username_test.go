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
			t.Errorf("expected the username to be valid, username: %v", username)
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
			t.Errorf("expected an error but got none, username: %v", username)
		}
	}
}

func TestUsername_ComputeUsernameFingerprint_1(t *testing.T) {
	username := Username("aa")
	fingerprint := ComputeUsernameFingerprint(username)
	fingerprintExpected := "961b6dd3ede3cb8ecbaacbd68de040cd78eb2ed5889130cceb4c49268ea4d506"

	if string(fingerprint) != fingerprintExpected {
		t.Errorf("got wrong fingerprint for username")
	}
}

func TestUsername_ComputeUsernameFingerprint_2(t *testing.T) {
	username := Username("aa")
	fingerprint1 := ComputeUsernameFingerprint(username)
	fingerprint2 := ComputeUsernameFingerprint(username)

	if string(fingerprint1) != string(fingerprint2) {
		t.Errorf("got different fingerprints for the same username")
	}
}
