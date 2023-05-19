package types

import (
	"github.com/google/uuid"
	"testing"
)

func TestId_NewId(t *testing.T) {
	id := NewId()

	_, err := uuid.Parse(string(id))

	if err != nil {
		t.Errorf("Expected valid UUID, but got error: %v", err)
	}
}

func TestId_ParseAndValidateId_Valid(t *testing.T) {
	validIds := []string{
		"e6f82bb7-a3f9-45c2-9a7c-4b79a6a46912",
		"e6f82bb7a3f945c29a7c4b79a6a46912",
		string(NewId()),
	}

	for _, id := range validIds {
		_, err := ParseAndValidateId(id)
		if err != nil {
			t.Error("Expected", id)
		}
	}
}

func TestId_ParseAndValidateId_Invalid(t *testing.T) {
	invalidIds := []string{
		"",
		"123",
		"notauuid",
	}

	for _, id := range invalidIds {
		_, err := ParseAndValidateId(id)
		if err == nil {
			t.Errorf("Expected error for id %v, but got none", id)
		}
	}
}
