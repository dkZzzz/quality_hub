package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// 创建一个临时的配置文件用于测试
	tempConfig := `{
		"sonarHost": "example.com",
		"sonarUser": "admin",
		"sonarPassword": "password",
		"mysqlHost": "localhost",
		"mysqlUser": "root",
		"mysqlPassword": "root",
		"mysqlDatabase": "test_db",
		"mysqlPort": "3306",
		"redisHost": "localhost",
		"redisPort": "6379",
		"etcdHost": "localhost",
		"etcdPort": "2379",
		"openaiSK": "secret_key",
		"codeStorePath": "/path/to/code",
		"userServerHost": "user.example.com",
		"sonarqubeServerHost": "sonar.example.com",
		"chatServerHost": "chat.example.com",
		"noticeServerHost": "notice.example.com"
	}`

	tempFile, err := os.CreateTemp("", "test_config.json")
	if err != nil {
		t.Fatalf("Failed to create temporary config file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(tempConfig)
	if err != nil {
		t.Fatalf("Failed to write to temporary config file: %v", err)
	}
	tempFile.Close()

	// 测试加载配置文件
	loadedConfig, err := LoadConfig(tempFile.Name())
	if err != nil {
		t.Errorf("LoadConfig() returned an error: %v", err)
	}

	// 检查加载的配置是否正确
	expectedConfig := &Config{
		SonarHost:           "example.com",
		SonarUser:           "admin",
		SonarPassword:       "password",
		MysqlHost:           "localhost",
		MysqlUser:           "root",
		MysqlPassword:       "root",
		MysqlDatabase:       "test_db",
		MysqlPort:           "3306",
		RedisHost:           "localhost",
		RedisPort:           "6379",
		EtcdHost:            "localhost",
		EtcdPort:            "2379",
		OpenaiSK:            "secret_key",
		CodeStorePath:       "/path/to/code",
		UserServerHost:      "user.example.com",
		SonarqubeServerHost: "sonar.example.com",
		ChatServerHost:      "chat.example.com",
		NoticeServerHost:    "notice.example.com",
	}

	if loadedConfig.SonarHost != expectedConfig.SonarHost ||
		loadedConfig.SonarUser != expectedConfig.SonarUser ||
		loadedConfig.SonarPassword != expectedConfig.SonarPassword {
		t.Error("Loaded config does not match expected config")
	}

	t.Logf("LoadConfig() test passed successfully")
}
