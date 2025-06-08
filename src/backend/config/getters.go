package config

import (
	"Melodex/src/logger"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var loadedConfig = GetConfig()

// GetConfig loads the configuration file.
// If loading fails, it initializes and saves a default configuration.
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

// ConfGetVolume retrieves the volume setting from the loaded configuration.
// If the format is invalid, it defaults to 100%.
func ConfGetVolume() int {
	volume, err := strconv.Atoi(strings.TrimSuffix(loadedConfig.General.DefaultVolume, "%"))
	if err != nil {
		logger.Log.Error("Invalid volume format in config, using default (100%)")
		return 100
	}
	logger.Log.Warnf("Using volume %d", volume)
	return volume
}

// ConfGetUser retrieves the configured user or falls back to the system user.
func ConfGetUser() string {
	user := loadedConfig.General.User
	if user == "" {
		user = os.Getenv("USER")
		logger.Log.Infof("User not valid, using environment user: %s", user)
	}
	logger.Log.Infof("Using user from config: %s", user)
	return user
}

// ConfGetDefaultPlaylist retrieves the default playlist path.
// Logs an error if no playlist is found.
func ConfGetDefaultPlaylist() string {
	playlist := loadedConfig.General.DefaultPlaylist
	if playlist == "" {
		logger.Log.Errorf("Playlist not found.")
		return ""
	}
	return playlist
}

// GetPlaylistArray retrieves the list of playlists from the configuration.
// Logs an error if no playlists are set.
func GetPlaylistArray() []string {
	playlist := loadedConfig.Playlist.Playlists
	if playlist == nil {
		logger.Log.Errorf("Playlist empty")
		return nil
	}
	return playlist
}

// GetBorderColor retrieves the configured border color.
// Returns an empty string if not defined.
func GetBorderColor() string {
	border := loadedConfig.Design.Border
	if border == "" {
		logger.Log.Error("Undefined border")
		return "#FFFFFF" // TODO Default Color
	}
	return border
}

// GetForegroundColor retrieves the configured foreground color.
// Returns an empty string if not defined.
func GetForegroundColor() string {
	fg := loadedConfig.Design.Foreground
	if fg == "" {
		logger.Log.Error("Undefined foreground color")
		return "#FFFFFF"
	}
	return fg
}

// GetBackgroundColor retrieves the configured background color.
// Returns an empty string if not defined.
func GetBackgroundColor() string {
	bg := loadedConfig.Design.Background
	if bg == "" {
		logger.Log.Error("Undefined background color")
		return "#000000"
	}
	return bg
}
