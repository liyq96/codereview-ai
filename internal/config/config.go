package config

import (
	"time"
)

type ServerConfig struct {
	Host         string        `yaml:"host"`
	Port         int           `yaml:"port"`
	Mode         string        `yaml:"mode"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type AiConfig struct {
	BaseUrl        string `yaml:"base_url"`
	ApiKey         string `yaml:"api_key"`
	ModelId        string `yaml:"model_id"`
	PromptFilePath string `yaml:"prompt_file_path"`
}

type SystemConfig struct {
	Server ServerConfig `yaml:"server"`
	Ai     AiConfig     `yaml:"ai"`
}
