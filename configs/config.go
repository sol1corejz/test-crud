package configs

import (
	"gopkg.in/yaml.v3"
	"os"
)

// Config структура для загрузки YAML
type Config struct {
	Server struct {
		Address string `yaml:"address"`
	} `yaml:"server"`

	Storage struct {
		Type             string `yaml:"type"`
		ConnectionString string `yaml:"connection_string"`
	} `yaml:"storage"`
}

// LoadConfig загружает конфигурацию из файла
func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
