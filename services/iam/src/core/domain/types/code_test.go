package types

import (
	"testing"
)

func TestCode_NewCode_ConcurrentCreation(t *testing.T) {
	nCodes := 1000

	codeMap := make(map[Code]bool)

	codes := make(chan Code, nCodes)
	errors := make(chan error, nCodes)

	for i := 0; i < nCodes; i++ {
		go func() {
			code, err := NewCode()
			if err != nil {
				errors <- err
			} else {
				codes <- code
			}
		}()
	}

	for i := 0; i < nCodes; i++ {
		select {
		case err := <-errors:
			t.Errorf("Expected no error, got %v", err)
		case code := <-codes:
			if _, exists := codeMap[code]; exists {
				t.Error("Expected no duplicate code", code)
			}
			if _, err := ParseAndValidateCode(string(code)); err != nil {
				t.Errorf("Expected a valid code, got %v", code)
			}
			codeMap[code] = true
		}
	}
}

// TestCode_ParseAndValidateCode_Valid tests that valid codes are parsed and validated correctly
func TestCode_ParseAndValidateCode_Valid(t *testing.T) {
	validCode := "IIr5d0IEBdLlVKkKNB97NtWaFKH11RH0nQWZ1dR46/s="
	_, err := ParseAndValidateCode(validCode)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}

// TestCode_ParseAndValidateCode_Invalid tests that giving invalid codes returns an error
func TestCode_ParseAndValidateCode_Invalid(t *testing.T) {
	var (
		tooShort        = "IIr5d0IEBdLlVKkKNB97NtWaFKH11RH0nQWZ1dR4/s="
		unsupportedChar = "{Ir5d0IEBdLlVKkKNB97NtWaFKH11RH0nQWZ1dR46/s="
	)

	invalidCodes := []string{tooShort, unsupportedChar}

	for c, invalidCode := range invalidCodes {
		p, err := ParseAndValidateCode(invalidCode)
		if err == nil {
			t.Errorf("Expected an error, got %v for %v", p, c)
		}
	}
}

// TestCode_EncodeForURL tests that valid codes are encoded correctly to be used in urls
func TestCode_EncodeForURL(t *testing.T) {
	codes := []struct {
		code     Code
		expected string
	}{
		{
			code:     Code("IIr5d0IEBdLlVKkKNB97NtWaFKH11RH0nQWZ1dR46/s="),
			expected: "SUlyNWQwSUVCZExsVktrS05COTdOdFdhRktIMTFSSDBuUVdaMWRSNDYvcz0=",
		},
		{
			code:     Code("NSfv0bK7Ewcm4+YAtE7JnRHTt7XDTP7RbUuQ22Ggzl8="),
			expected: "TlNmdjBiSzdFd2NtNCtZQXRFN0puUkhUdDdYRFRQN1JiVXVRMjJHZ3psOD0=",
		},
	}

	for _, c := range codes {
		got := c.code.EncodeForURL()

		if got != c.expected {
			t.Errorf("Expected %s, got %s", c.expected, got)
		}
	}
}
