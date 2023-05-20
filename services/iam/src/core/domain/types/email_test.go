package types

import (
	"strings"
	"testing"
)

func TestEmail_ParseAndValidateEmail_Valid(t *testing.T) {
	validEmails := []string{
		"abc@gmail.com",
		"abc.def-ghi@gmail.com",
		"aBc@gmail.com",
		" aBc@gmail.com ",
	}

	for _, validEmail := range validEmails {
		email, err := ParseAndValidateEmail(validEmail)

		if err != nil {
			t.Errorf("Ecpected no error, got %s", err)
		}

		expected := strings.ToLower(strings.TrimSpace(validEmail))

		if string(email) != expected {
			t.Errorf("Expected %s, got %s", expected, email)
		}
	}
}

func Test_ParseAndValidateEmail_Invalid(t *testing.T) {
	invalidEmails := []string{
		"abc",
		"ab(c@gmail.com",
		"",
	}

	for _, invalidEmail := range invalidEmails {
		email, err := ParseAndValidateEmail(invalidEmail)

		if err == nil {
			t.Errorf("Expected %s, got %s", invalidEmail, email)
		}
	}
}
