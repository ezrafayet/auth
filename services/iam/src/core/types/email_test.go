package types

import (
	"testing"
)

func Test_ParseAndValidateEmail_AcceptValidEmail1(t *testing.T) {
	originalEmail := "abc@gmail.com"
	email, err := ParseAndValidateEmail(originalEmail)
	if err != nil {
		t.Error("Expected to accept valid email")
	}
	if string(email) != originalEmail {
		t.Error("Email value is not expected")
	}
}

func Test_ParseAndValidateEmail_AcceptValidEmail2(t *testing.T) {
	originalEmail := "abc.def-ghi@gmail.com"
	email, err := ParseAndValidateEmail(originalEmail)
	if err != nil {
		t.Error("Expected to accept valid email")
	}
	if string(email) != originalEmail {
		t.Error("Email value is not expected")
	}
}

func Test_ParseAndValidateEmail_ReformatAndAcceptEmail1(t *testing.T) {
	originalEmail := " abc@gmail.com "
	email, err := ParseAndValidateEmail(originalEmail)
	if err != nil {
		t.Error("Expected to accept valid email")
	}
	if string(email) != "abc@gmail.com" {
		t.Error("Email value is not expected")
	}
}

func Test_ParseAndValidateEmail_ReformatAndAcceptEmail2(t *testing.T) {
	originalEmail := " aBc@gmail.com "
	email, err := ParseAndValidateEmail(originalEmail)
	if err != nil {
		t.Error("Expected to accept valid email")
	}
	if string(email) != "abc@gmail.com" {
		t.Error("Email value is not expected")
	}
}

func Test_ParseAndValidateEmail_RejectInvalidEmail1(t *testing.T) {
	originalEmail := "abc"
	_, err := ParseAndValidateEmail(originalEmail)
	if err == nil {
		t.Error("Expected to accept valid email")
	}
}

func Test_ParseAndValidateEmail_RejectInvalidEmail2(t *testing.T) {
	originalEmail := "abcd(abc@gmail.com"
	_, err := ParseAndValidateEmail(originalEmail)
	if err == nil {
		t.Error("Expected to accept valid email")
	}
}

func Test_ParseAndValidateEmail_RejectInvalidEmail3(t *testing.T) {
	originalEmail := ""
	_, err := ParseAndValidateEmail(originalEmail)
	if err == nil {
		t.Error("Expected to accept valid email")
	}
}
