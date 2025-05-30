package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

const ConfigFile = "config.toml"

var ConfigDir = filepath.Join(os.Getenv("HOME"), ".config/melodex")

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

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, logError("reading config file", err)
	}

	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, logError("parsing config", err)
	}

	return &cfg, nil
}

func SaveConfig(path string, cfg *Config) error {
	data, err := toml.Marshal(cfg)
	if err != nil {
		return logError("encoding config", err)
	}

	return os.WriteFile(path, data, 0644)
}

func DefaultConfig() *Config {
	user := getUser()
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

func getUser() string {
	if user := os.Getenv("USER"); user != "" {
		return user
	}
	return "unknown"
}

func GenerateDefaultConfig() bool {
	configPath := filepath.Join(ConfigDir, ConfigFile)

	if err := SaveConfig(configPath, DefaultConfig()); err != nil {
		logError("generating default config", err)
		return false
	}

	log.Println("Default config generated successfully.")
	return true
}

func logError(action string, err error) error {
	log.Printf("Error %s: %v\n", action, err)
	return err
}
