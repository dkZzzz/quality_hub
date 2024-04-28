package mysql

import (
	"testing"
)

func TestPasswordEncryption(t *testing.T) {
	// 设置测试用的密码和参数
	password := "test_password"
	argon2Params := &Argon2Params{
		Memory:      65536,
		Iterations:  1,
		Parallelism: 2,
		KeyLength:   32,
		SaltLength:  16,
	}

	// 测试密码加密
	encodedHash, err := generateFromPassword(password, argon2Params)
	if err != nil {
		t.Errorf("generateFromPassword() returned an error: %v", err)
	}

	if encodedHash == "" {
		t.Error("Generated encoded hash is empty")
	}

	// 测试密码验证
	match, err := comparePasswordAndHash(password, encodedHash)
	if err != nil {
		t.Errorf("comparePasswordAndHash() returned an error: %v", err)
	}

	if !match {
		t.Error("Passwords did not match")
	}

	t.Logf("Password encryption and verification test passed successfully")
}
