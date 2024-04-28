package clone

import (
	"os"
	"testing"

	"github.com/dkZzzz/quality_hub/config"
)

func TestClone(t *testing.T) {
	// 设置测试用的 URL 和名称
	url := "https://github.com/example/repository.git"
	name := "test_repository"

	// 运行 Clone 函数
	err := Clone(url, name)

	// 检查错误
	if err != nil {
		t.Errorf("Clone() returned an error: %v", err)
	}

	// 检查目标目录是否存在
	targetDir := config.Cfg.CodeStorePath + name
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		t.Errorf("Target directory does not exist after cloning: %s", targetDir)
	}

	// 如果一切正常，打印测试通过的消息
	t.Logf("Clone() test passed successfully")
}
