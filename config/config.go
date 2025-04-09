package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	MySQL struct {
		DataSource string `yaml:"DataSource"`
	} `yaml:"MySQL"`
	Redis struct {
		Host string `yaml:"Host"`
		Pass string `yaml:"Pass"`
	} `yaml:"Redis"`
	Kafka struct {
		Brokers []string `yaml:"Brokers"`
		Topic   string   `yaml:"Topic"`
	} `yaml:"Kafka"`
}

func LoadConfig(filePath string, cfg *Config) error {
	// LoadConfig 逻辑
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return fmt.Errorf("无法重新解组配置文件: %w", err)
	}

	return nil
}
