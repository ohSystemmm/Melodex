package config

import (
	"Melodex/src/log"
	"Melodex/src/settings"
	"github.com/pelletier/go-toml/v2"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// LoadConfig reads a configuration file and returns the parsed Config struct.
func LoadConfig(cfgFile string) (Config, error) {
	if cfgFile == "" {
		cfgFile = filepath.Join(settings.AppConfig.ConfigDir, settings.AppConfig.ConfigFile)
	}

	data, err := os.ReadFile(cfgFile)
	if err != nil {
		log.Log.Warnf("Config file not found: %s. Falling back to default.", cfgFile)
		return GenerateDefaultConfig(), err
	}

	var cfg Config
	if err = toml.Unmarshal(data, &cfg); err != nil {
		log.Log.Warnf("Error parsing config file: %s. Using default settings.", cfgFile)
		return GenerateDefaultConfig(), err
	}
	return cfg, nil
}

// GenerateDefaultConfig initializes default values while ensuring portability.
func GenerateDefaultConfig() Config {
	user := settings.GetUser()
	homePath, err := os.UserHomeDir()
	if err != nil {
		homePath = filepath.Join("/home", user) // Fallback if UserHomeDir fails
	}

	return Config{
		General: General{
			User:            user,
			SetUp:           "done",
			DefaultPlaylist: filepath.Join(homePath, "Music"),
			DefaultVolume:   "75%",
		},
		Playlist: Playlist{
			Playlists: generateDefaultPlaylists(homePath),
		},
		Design: Design{
			BorderColor:      "#FFFFFF",
			ForegroundColor:  "#FFFFFF",
			BackgroundColor:  "#000000",
			VolumeBarColor:   "#00FFFF",
			MusicSliderColor: "#FF00FF",
		},
	}
}

// Helper function to generate default playlists dynamically.
func generateDefaultPlaylists(homePath string) []string {
	return []string{
		filepath.Join(homePath, "example_playlist1"),
		filepath.Join(homePath, "example_playlist2"),
		filepath.Join(homePath, "example_playlist3"),
	}
}

// Global configuration variable.
var config = setCFG()

// LoadGlobalConfig initializes the global configuration safely.
func setCFG() Config {
	cfg, err := LoadConfig(settings.AppConfig.ConfigDir + settings.AppConfig.ConfigFile)
	if err != nil {
		log.Log.Errorf("Config file not found: %s. Using default settings.", settings.AppConfig.ConfigFile)
		return GenerateDefaultConfig()
	}
	return cfg
}

// Accessor functions
func GetUser() string  { return config.General.User }
func GetSetUp() string { return config.General.SetUp }
func GetDefaultVolume() int {
	volumeStr := config.General.DefaultVolume

	// Remove the '%' character if present
	volumeStr = strings.TrimSuffix(volumeStr, "%")

	volume, err := strconv.Atoi(volumeStr)
	if err != nil {
		log.Log.Errorf("Error converting volume (%s) to integer: %v", volumeStr, err)
		return 0 // Fallback to safe default
	}

	return volume
}

func GetDefaultPlaylist() string { return config.General.DefaultPlaylist }
func GetPlaylists() []string     { return config.Playlist.Playlists }
func GetBorderColor() string     { return config.Design.BorderColor }
func GetForegroundColor() string { return config.Design.ForegroundColor }
func GetBackgroundColor() string { return config.Design.BackgroundColor }
