package types

import (
	"testing"
)

// TestAuthMethod_String_Password tests that the code 0 is converted to "password"
func TestAuthMethod_String_Password(t *testing.T) {
	var a AuthType = 1
	expected := "password"
	result := a.String()
	if result != expected {
		t.Errorf("Expected %v, got %v", "password", result)
	}
}

// TestAuthMethod_String_MagicLink tests that the code 1 is converted to "magic-link"
func TestAuthMethod_String_MagicLink(t *testing.T) {
	var a AuthType = 2
	expected := "magic-link"
	result := a.String()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestAuthMethod_String_Invalid tests that an invalid code is converted to "unknown"
func TestAuthMethod_String_Invalid(t *testing.T) {
	var a AuthType = 0
	expected := "unknown"
	result := a.String()
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestAuthMethod_ParseAndValidate_Password tests that "password" is converted to 0
func TestAuthMethod_ParseAndValidate_Password(t *testing.T) {
	m, err := ParseAndValidateAuthType("password")
	expectedM := 1
	expectedErr := error(nil)
	if int(m) != expectedM || err != expectedErr {
		t.Errorf("Expected (%v, %v) got (%v, %v)", expectedM, expectedErr, m, err)
	}
}

// TestAuthMethod_ParseAndValidate_MagicLink tests that "magic-link" is converted to 1
func TestAuthMethod_ParseAndValidate_MagicLink(t *testing.T) {
	m, err := ParseAndValidateAuthType("magic-link")
	expectedM := 2
	expectedErr := error(nil)
	if int(m) != expectedM || err != expectedErr {
		t.Errorf("Expected (%v, %v) got (%v, %v)", expectedM, expectedErr, m, err)
	}
}

// TestAuthMethod_ParseAndValidate_Invalid tests that passing an invalid string gives an error
func TestAuthMethod_ParseAndValidate_Invalid(t *testing.T) {
	_, err := ParseAndValidateAuthType("foobar")
	if err == nil {
		t.Errorf("Expected an error")
	}
}
