package redis

import (
	"context"
	"testing"
)

func TestJWTMatch(t *testing.T) {
	// 假设你有一个测试的用户名和 JWT token
	username := "test_user"
	token := "test_token"

	// 测试 JWTMatch 函数
	match := JWTMatch(context.Background(), username, token)

	if match {
		t.Logf("JWT token matches for user %s", username)
	} else {
		t.Errorf("JWT token does not match for user %s", username)
	}
}
