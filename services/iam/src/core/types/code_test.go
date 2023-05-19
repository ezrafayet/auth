package types

import (
	"testing"
)

func Test_Code_NewCode_ConcurrentCreation(t *testing.T) {
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
			t.Errorf("Error generating code: %v", err)
		case code := <-codes:
			if _, exists := codeMap[code]; exists {
				t.Errorf("Duplicate code found: %v", code)
			}
			if _, err := ParseAndValidateCode(string(code)); err != nil {
				t.Errorf("Invalid code found: %v", code)
			}
			codeMap[code] = true
		}
	}
}

func Test_Code_ParseAndValidateCode_Valid(t *testing.T) {
	validCode := "IIr5d0IEBdLlVKkKNB97NtWaFKH11RH0nQWZ1dR46/s="
	code, err := ParseAndValidateCode(validCode)
	if code != Code(validCode) {
		t.Errorf("Wrong code parsing, got %s", code)
	}
	if err != nil {
		t.Errorf("Wrong code parsing, got error %s", err)
	}
}

func Test_Code_ParseAndValidateCode_InvalidLength(t *testing.T) {
	validCode := "IIr5d0IEBdLlVKkKNB97NtWaFKH11RH0nQWZ1dR4/s="
	_, err := ParseAndValidateCode(validCode)
	if err == nil {
		t.Errorf("Wrong code parsing, got error %s", err)
	}
}

func Test_Code_ParseAndValidateCode_InvalidChar(t *testing.T) {
	validCode := "{Ir5d0IEBdLlVKkKNB97NtWaFKH11RH0nQWZ1dR46/s="
	_, err := ParseAndValidateCode(validCode)
	if err == nil {
		t.Errorf("Wrong code parsing, got error %s", err)
	}
}

func Test_Code_EncodeForURL(t *testing.T) {
	c1 := Code("IIr5d0IEBdLlVKkKNB97NtWaFKH11RH0nQWZ1dR46/s=")

	if c1.EncodeForURL() != "SUlyNWQwSUVCZExsVktrS05COTdOdFdhRktIMTFSSDBuUVdaMWRSNDYvcz0=" {
		t.Errorf("Wrong code encoding for url, got %s", c1.EncodeForURL())
	}

	c2 := Code("NSfv0bK7Ewcm4+YAtE7JnRHTt7XDTP7RbUuQ22Ggzl8=")

	if c2.EncodeForURL() != "TlNmdjBiSzdFd2NtNCtZQXRFN0puUkhUdDdYRFRQN1JiVXVRMjJHZ3psOD0=" {
		t.Errorf("Wrong code encoding for url, got %s", c1.EncodeForURL())
	}
}
