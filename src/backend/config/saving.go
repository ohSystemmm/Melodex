package config

import (
	"Melodex/src/log"
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"os"
	"path/filepath"
)

func SaveConfig(path string, cfg *Config) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := toml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to serialize config data: %w", err)
	}

	if err = os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	log.Log.Info("Config file saved successfully")
	return nil
}
