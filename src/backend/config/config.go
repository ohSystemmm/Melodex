package config

import (
	"Melodex/src/logger"
	"github.com/pelletier/go-toml/v2"
	"os"
	"path/filepath"
)

var (
	fileConfig = "config.toml"
	dirConfig  = filepath.Join(os.Getenv("HOME"), ".config/melodex")
)

type Config struct {
	General  General  `toml:"general"`
	Playlist Playlist `toml:"playlist"`
	Design   Design   `toml:"design"`
}
type General struct {
	User            string `toml:"user"`
	DefaultPlaylist string `toml:"default_playlist"`
	DefaultVolume   string `toml:"default_volume"`
}
type Playlist struct {
	Playlists []string `toml:"playlists"`
}
type Design struct {
	Border     string `toml:"border"`
	Foreground string `toml:"foreground"`
	Background string `toml:"background"`
}

// FUNCTION: Tries loading the config
func LoadConfig(path string) *Config {
	data, err := os.ReadFile(path)
	if err != nil {
		logger.Log.Error("Load config file error:", err)
		return nil
	}

	var cfg Config
	if err = toml.Unmarshal(data, &cfg); err != nil {
		logger.Log.Error("Load config file error:", err)
		return nil
	}

	return &cfg
}

// FUNCTION: Saves the config
func SaveConfig(path string, cfg *Config) bool {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			logger.Log.Error("Failed to create config directory:", err)
			return false
		}
	}

	data, err := toml.Marshal(cfg)
	if err != nil {
		logger.Log.Error("Save config file error:", err)
		return false
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		logger.Log.Error("Save config file error:", err)
		return false
	}

	return true
}

func getEnvUser() string {
	user := os.Getenv("USER")
	if user == "" {
		return "unknown"
	}
	return user
}

// FUNCTION: Template for the default config
func DefaultConfig() *Config {
	user := getEnvUser()
	basePath := filepath.Join("/home", user)

	return &Config{
		General: General{
			User:            user,
			DefaultPlaylist: filepath.Join(basePath, "example_playlist"),
			DefaultVolume:   "100%",
		},
		Playlist: Playlist{
			Playlists: []string{
				filepath.Join(basePath, "example_playlist1"),
				filepath.Join(basePath, "example_playlist2"),
				filepath.Join(basePath, "example_playlist3"),
			},
		},
		Design: Design{
			Border:     "#FFFFFF",
			Foreground: "#FFFFFF",
			Background: "#000000",
		},
	}
}

// FUNCTION: Generates the default config
func GenerateDefaultConfig() bool {
	configPath := filepath.Join(dirConfig, fileConfig)

	if !SaveConfig(configPath, DefaultConfig()) {
		logger.Log.Error("Generate default configuration failed")
		return false
	}
	logger.Log.Info("Generate default configuration success")
	return true
}
