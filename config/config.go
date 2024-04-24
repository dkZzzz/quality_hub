package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

var (
	Cfg *Config
)

type Config struct {
	SonarHost     string `json:"sonarHost"`
	SonarUser     string `json:"sonarUser"`
	SonarPassword string `json:"sonarPassword"`

	MysqlHost     string `json:"mysqlHost"`
	MysqlUser     string `json:"mysqlUser"`
	MysqlPassword string `json:"mysqlPassword"`
	MysqlDatabase string `json:"mysqlDatabase"`
	MysqlPort     string `json:"mysqlPort"`

	RedisHost string `json:"redisHost"`
	RedisPort string `json:"redisPort"`

	EtcdHost string `json:"etcdHost"`
	EtcdPort string `json:"etcdPort"`

	OpenaiSK string `json:"openaiSK"`

	CodeStorePath string `json:"codeStorePath"`

	UserServerHost      string `json:"userServerHost"`
	SonarqubeServerHost string `json:"sonarqubeServerHost"`
	ChatServerHost      string `json:"chatServerHost"`
	NoticeServerHost    string `json:"noticeServerHost"`
}

func init() {
	_, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	log.Println(Cfg)
}

func LoadConfig() (*Config, error) {
	config := &Config{}
	file, err := os.Open("config.json")

	if err != nil {
		return config, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	Cfg = config

	return config, nil
}
