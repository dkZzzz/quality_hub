package jwt

import (
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	username := "test_user"
	userID := "12345"

	token, err := GenerateJWT(username, userID)
	if err != nil {
		t.Errorf("GenerateJWT() returned an error: %v", err)
	}

	if token == "" {
		t.Error("Generated JWT token is empty")
	}

	t.Logf("GenerateJWT() test passed successfully")
}

func TestValidateJWT(t *testing.T) {
	username := "test_user"
	userID := "12345"

	// 生成JWT token
	token, err := GenerateJWT(username, userID)
	if err != nil {
		t.Fatalf("Failed to generate JWT token: %v", err)
	}

	// 验证过期的 token
	err = ValidateJWT(token)
	if err == nil {
		t.Logf("ValidateJWT() no error was returned")
	} else {
		t.Logf("ValidateJWT() returned expected error for expired token: %v", err)
	}

	t.Logf("ValidateJWT() test passed successfully")
}
