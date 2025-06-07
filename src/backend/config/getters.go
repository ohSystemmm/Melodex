package config

import (
	"Melodex/src/logger"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var loadedConfig = GetConfig()

func GetConfig() *Config {
	configPath := filepath.Join(dirConfig, fileConfig)
	cfg := LoadConfig(configPath)

	if cfg == nil {
		logger.Log.Errorf("Error loading config: %s", configPath)

		defaultCfg := DefaultConfig()
		logger.Log.Infof("Using default config: %+v", defaultCfg)

		saveErr := SaveConfig(configPath, defaultCfg)
		if !saveErr {
			logger.Log.Errorf("Failed to save default config")
			return defaultCfg
		}

		return defaultCfg
	}
	return cfg
}

func ConfGetVolume() int {
	volume, err := strconv.Atoi(strings.TrimSuffix(loadedConfig.General.DefaultVolume, "%"))
	if err != nil {
		logger.Log.Error("Invalid volume format in config, using default (100%)")
		return 100
	}
	logger.Log.Warnf("Using volume %d", volume)
	return volume
}
func ConfGetUser() string {
	user := loadedConfig.General.User
	if user == "" {
		user = os.Getenv("USER")
		logger.Log.Infof("User not valid, using environment user: %s", user)
	}
	logger.Log.Infof("Using user from config: %s", user)
	return user
}
func ConfGetDefaultPlaylist() string {
	playlist := loadedConfig.General.DefaultPlaylist
	if playlist == "" {
		logger.Log.Errorf("Playlist not found.")
		return ""
	}
	return playlist
}

func GetPlaylistArray() []string {
	playlist := loadedConfig.Playlist.Playlists
	if playlist == nil {
		logger.Log.Errorf("Playlist empty")
		return nil
	}
	return playlist
}

func GetBorderColor() string {
	border := loadedConfig.Design.Border
	if border == "" {
		logger.Log.Error("Undefined border")
		return "" // TODO Default Color
	}
	return border
}

func GetForegroundColor() string {
	fg := loadedConfig.Design.Foreground
	if fg == "" {
		logger.Log.Error("Undefined foreground color")
		return "" // TODO Default Color
	}
	return fg
}

func GetBackgroundColor() string {
	bg := loadedConfig.Design.Background
	if bg == "" {
		logger.Log.Error("Undefined background color")
		return "" // TODO Default Color
	}
	return bg
}
