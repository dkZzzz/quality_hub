package mysql

import (
	"testing"
)

func TestInitDB(t *testing.T) {
	// 假设你已经正确配置了 MySQL 数据库，并且能够连接到数据库
	_, err := InitDB()
	if err != nil {
		t.Errorf("InitDB() returned an error: %v", err)
	}

	// 这里可以添加更多的测试，比如查询数据库是否正常等

	t.Logf("InitDB() test passed successfully")
}
