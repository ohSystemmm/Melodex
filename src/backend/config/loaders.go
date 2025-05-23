package config

import (
	"log"
	"strconv"
)

func GetConfig() (*Config, string, error) {
	cfg, err := LoadConfig(ConfigDir + ConfigFile)
	if err != nil {
		log.Printf("Error loading config: %v\n", err)
		errMsg := "Failed to load config file. "

		defaultCfg := DefaultConfig()

		saveErr := SaveConfig(ConfigDir+ConfigFile, defaultCfg)
		if saveErr != nil {
			log.Printf("Failed to save default config: %v\n", saveErr)

			return defaultCfg, "", saveErr
		}

		return defaultCfg, errMsg, nil
	}
	return cfg, "", nil
}

func GetDefaultPlaylist() string {
	cfg, _, err := GetConfig()
	if err != nil {
		return ""
	}
	return cfg.General.DefaultPlaylist
}

func GetVolume() int {
	cfg, _, err := GetConfig()
	if err != nil {
		return 100
	}

	volume, err := strconv.Atoi(cfg.General.DefaultVolume[:len(cfg.General.DefaultVolume)-1])
	if err != nil {
		return 100
	}

	return volume
}

func GetPlaylists() []string {
	cfg, _, err := GetConfig()
	if err != nil {
		return nil
	}
	return cfg.Playlist.Playlists
}

func GetBorder() string {
	cfg, _, err := GetConfig()
	if err != nil {
		return ""
	}
	return cfg.Design.Border
}

func GetForeground() string {
	cfg, _, err := GetConfig()
	if err != nil {
		return ""
	}
	return cfg.Design.Foreground
}

func GetBackground() string {
	cfg, _, err := GetConfig()
	if err != nil {
		return ""
	}
	return cfg.Design.Background
}
