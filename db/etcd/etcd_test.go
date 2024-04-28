package etcd

import (
	"testing"
)

func TestGet(t *testing.T) {
	service := "test_service"

	// 测试 Get 函数
	exists, err := Get(service)
	if err != nil {
		t.Errorf("Get() returned an error: %v", err)
	}

	if exists {
		t.Logf("Service %s exists in etcd", service)
	} else {
		t.Logf("Service %s does not exist in etcd", service)
	}
}

func TestRegister(t *testing.T) {
	service := "test_service"

	// 测试 Register 函数
	err := Register(service)
	if err != nil {
		t.Errorf("Register() returned an error: %v", err)
	} else {
		t.Logf("Registered service %s successfully", service)
	}
}
