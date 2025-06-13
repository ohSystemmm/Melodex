package config

import (
	"Melodex/src/log"
	"github.com/pelletier/go-toml/v2"
	"os"
)

func LoadConfig() *Config {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Log.Errorf("Error loading config file (%s): %v", configFilePath, err)
		return DefaultConfig(true)
	}

	var cfg Config
	if err = toml.Unmarshal(data, &cfg); err != nil {
		log.Log.Errorf("Error parsing config file (%s): %v", configFilePath, err)
		return DefaultConfig(false)
	}

	return &cfg
}
