package config

import (
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

const ConfigFile = "config.toml"

var ConfigDir = "/home/" + getUser() + "/.config/melodex/"

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
		log.Printf("Error reading config file: %v\n", err)
		return nil, err
	}

	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		log.Printf("Error parsing config: %v\n", err)
		return nil, err
	}

	return &cfg, nil
}

func SaveConfig(path string, cfg *Config) error {
	data, err := toml.Marshal(cfg)
	if err != nil {
		log.Printf("Error encoding config: %v\n", err)
		return err
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		log.Printf("Error writing config file: %v\n", err)
		return err
	}

	return nil
}

func DefaultConfig() *Config {
	user := getUser()
	return &Config{
		General: General{
			User:            user,
			DefaultPlaylist: "/home/" + user + "/example_playlist",
			DefaultVolume:   "100%",
		},
		Playlist: Playlist{
			Playlists: []string{
				"/home/" + user + "/example_playlist1",
				"/home/" + user + "/example_playlist2",
				"/home/" + user + "/example_playlist3",
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
	user := os.Getenv("USER")
	if user != "" {
		return user
	}
	return "unknown"
}

func GenerateDefaultConfig() bool {
	defaultConfig := DefaultConfig()
	configPath := ConfigDir + ConfigFile

	err := SaveConfig(configPath, defaultConfig)
	if err != nil {
		log.Printf("Failed to generate default config: %v\n", err)
		return false
	}

	log.Println("Default config generated successfully.")
	return true
}
