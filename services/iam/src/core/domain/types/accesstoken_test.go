package types

import (
	"os"
	"testing"
	"time"
)

func TestAccessToken_Generate(t *testing.T) {
	os.Setenv("JWT_PRIVATE_KEY", "s3cr3t")

	customClaims := CustomClaims{
		UserId:       "user_id",
		Roles:        "user",
		ServerRegion: "us-east-1",
	}

	stringToken, _, err := NewAccessToken(customClaims, time.Now())

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if stringToken == "" {
		t.Errorf("Expected token, got empty string")
	}
}

func TestAccessToken_ParseAndValidate_ValidToken(t *testing.T) {
	os.Setenv("JWT_PRIVATE_KEY", "s3cr3t")

	customClaims := CustomClaims{
		UserId:       "user_id",
		Roles:        "user",
		ServerRegion: "us-east-1",
	}

	stringToken, _, _ := NewAccessToken(customClaims, time.Now())

	isValid, claims, err := ParseAndValidateAccessToken(string(stringToken))

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !isValid {
		t.Errorf("Expected valid token, got invalid")
	}

	if claims.UserId != customClaims.UserId {
		t.Errorf("Expected %v, got %v", customClaims.UserId, claims.UserId)
	}

	if claims.Roles != customClaims.Roles {
		t.Errorf("Expected %v, got %v", customClaims.Roles, claims.Roles)
	}

	if claims.ServerRegion != customClaims.ServerRegion {
		t.Errorf("Expected %v, got %v", customClaims.ServerRegion, claims.ServerRegion)
	}
}

func TestAccessToken_ParseAndValidate_ExpiredToken(t *testing.T) {
	os.Setenv("JWT_PRIVATE_KEY", "s3cr3t")

	customClaims := CustomClaims{
		UserId:       "user_id",
		Roles:        "user",
		ServerRegion: "us-east-1",
	}

	stringToken, _, _ := NewAccessToken(customClaims, time.Now().Add(time.Minute*-16))

	isValid, _, err := ParseAndValidateAccessToken(string(stringToken))

	if err == nil {
		t.Errorf("Expected error, got no nil")
	}

	if isValid {
		t.Errorf("Expected true, got false")
	}
}

func TestAccessToken_ParseAndValidate_InvalidToken(t *testing.T) {
	os.Setenv("JWT_PRIVATE_KEY", "s3cr3t")

	isValid, _, err := ParseAndValidateAccessToken("invalid_token")

	if err == nil {
		t.Errorf("Expected error, got no nil")
	}

	if isValid {
		t.Errorf("Expected true, got false")
	}
}
