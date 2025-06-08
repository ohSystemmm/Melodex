package config

import (
	"Melodex/src/log"
	"Melodex/src/settings"
	"github.com/pelletier/go-toml/v2"
	"os"
	"path/filepath"
)

func LoadConfig(cfgFile string) *Config {
	if cfgFile == "" {
		cfgFile = filepath.Join(settings.AppConfig.ConfigDir, settings.AppConfig.ConfigFile)
	}

	data, err := os.ReadFile(cfgFile)
	if err != nil {
		log.Log.Warn("Config file not found. Falling back to default.")
		GlobalConfig = GenerateDefaultConfig()
		return GlobalConfig
	}

	var cfg Config
	if err = toml.Unmarshal(data, &cfg); err != nil {
		log.Log.Warn("Error parsing config file. Using default settings.")
		GlobalConfig = GenerateDefaultConfig()
		return GlobalConfig
	}

	GlobalConfig = &cfg
	return GlobalConfig
}
