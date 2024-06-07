package config

import (
	"github.com/BurntSushi/toml"
)

type DatabaseConfig struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
}

type DebugConfig struct {
	Enabled bool
	Display bool
	Log     bool
}

type RouterConfig struct {
	Port      string
	AdminPath string
}

type Config struct {
	Database DatabaseConfig
	Debug    DebugConfig
	Router   RouterConfig
}

func ParseConfig() (*Config, error) {
	var config Config

	_, err := toml.DecodeFile("./config.toml", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
