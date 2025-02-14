package proxy

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ProxyPort string `yaml:"proxy_port"`
	TargetURL string `yaml:"target_url"`
	LogFile   string `yaml:"log_file"`
}

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
