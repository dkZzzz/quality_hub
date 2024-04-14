package clone

import (
	"os"

	"github.com/dkZzzz/quality_hub/config"
	"github.com/go-git/go-git/v5"
)

// 执行git clone命令
func Clone(url, name string) error {
	targetDir := config.Cfg.CodeStorePath + name
	_, err := git.PlainClone(targetDir, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}
	return nil
}
