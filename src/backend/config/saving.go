package config

import (
	"Melodex/src/log"
	"Melodex/src/settings"
	"github.com/pelletier/go-toml/v2"
	"os"
	"path/filepath"
)

func SaveConfig(cfg *Config) bool {
	configPath := settings.AppConfig.ConfigDir + settings.AppConfig.ConfigFile

	dir := filepath.Dir(configPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Log.Error("Failed to create config directory:", err)
			return false
		}
	}

	data, err := toml.Marshal(cfg)
	if err != nil {
		log.Log.Error("Error marshaling config file:", err)
		return false
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		log.Log.Error("Error writing config file:", err)
		return false
	}

	return true
}
