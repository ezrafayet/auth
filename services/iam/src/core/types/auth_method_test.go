package types

import (
	"testing"
)

func Test_AuthMethod_String_Password(t *testing.T) {
	var a AuthMethod = 0
	if a.String() != "password" {
		t.Errorf("Did not get the correct auth method, got %v", a.String())
	}
}

func Test_AuthMethod_String_MagicLink(t *testing.T) {
	var a AuthMethod = 1
	if a.String() != "magic-link" {
		t.Errorf("Did not get the correct auth method, got %v", a.String())
	}
}

func Test_AuthMethod_String_Invalid(t *testing.T) {
	var a AuthMethod = 127
	if a.String() != "unknown" {
		t.Errorf("Did not get the correct auth method, got %v", a.String())
	}
}

func Test_AuthMethod_ParseAndValidate_Password(t *testing.T) {
	m, err := ParseAndValidateAuthMethod("password")
	if m != 0 {
		t.Errorf("ParseAndValidateAuthMethod() = %v, want %v", m, AuthMethodPassword)
	}
	if err != nil {
		t.Errorf("ParseAndValidateAuthMethod() = %v, want %v", err, nil)
	}
}

func Test_ParseAndValidateAuthMethod_MagicLink(t *testing.T) {
	m, err := ParseAndValidateAuthMethod("magic-link")
	if m != 1 {
		t.Errorf("ParseAndValidateAuthMethod() = %v, want %v", m, AuthMethodPassword)
	}
	if err != nil {
		t.Errorf("ParseAndValidateAuthMethod() = %v, want %v", err, nil)
	}
}

func Test_ParseAndValidateAuthMethod_Invalid(t *testing.T) {
	_, err := ParseAndValidateAuthMethod("foobar")
	if err == nil {
		t.Errorf("ParseAndValidateAuthMethod() = %v, want %v", err, nil)
	}
}
