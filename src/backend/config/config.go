package config

import (
	"Melodex/src/logger"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

var (
	// fileConfig defines the name of the configuration file.
	fileConfig = "config.toml"

	// dirConfig specifies the directory where the configuration file is stored.
	dirConfig = filepath.Join(os.Getenv("HOME"), ".config/melodex")
)

// Config represents the main configuration structure, grouping General, Playlist, and Design settings.
type Config struct {
	General  General  `toml:"general"`
	Playlist Playlist `toml:"playlist"`
	Design   Design   `toml:"design"`
}

// General contains general application settings such as user information and default values.
type General struct {
	User            string `toml:"user"`             // User defines the current system user.
	DefaultPlaylist string `toml:"default_playlist"` // DefaultPlaylist specifies the default playlist path.
	DefaultVolume   string `toml:"default_volume"`   // DefaultVolume sets the default volume level.
}

// Playlist manages a collection of user-defined playlists.
type Playlist struct {
	Playlists []string `toml:"playlists"` // Playlists holds a list of paths to user-created playlists.
}

// Design holds the visual settings for the application's interface.
type Design struct {
	Border     string `toml:"border"`     // Border color setting for the UI.
	Foreground string `toml:"foreground"` // Foreground color setting for text and icons.
	Background string `toml:"background"` // Background color setting for the UI.
}

// getEnvUser retrieves the current system user from the environment variables.
// If the user is not found, it returns "unknown".
func getEnvUser() string {
	user := os.Getenv("USER")
	if user == "" {
		return "unknown"
	}
	return user
}

// LoadConfig reads the configuration file from the given path and un-marshals it into a Config struct.
// If an error occurs, it logs the error and returns nil.
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

// SaveConfig writes the provided configuration to the specified file path.
// It creates the directory if it does not exist. Returns true on success, false on failure.
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

// DefaultConfig initializes a default configuration with predefined values,
// using the system user to determine base file paths.
func DefaultConfig() *Config {
	user := getEnvUser()
	basePath := filepath.Join("/home", user)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Do you wish to set the default playlist? (this is the playlist that melodex enters per default) (y/n)")
	input, _ := reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))

	defaultPlaylist := filepath.Join(basePath, "example_playlist")

	switch input {
	case "y", "yes":
		fmt.Println("Enter full location from root e.g. /home/johndoe/music")
		_, err := fmt.Scanln(&defaultPlaylist)
		if err != nil {
			logger.Log.Error("Invalid Input in selecting default Playlist")
		}
	case "n", "no":
		fmt.Println("Choosing default option: " + defaultPlaylist)
	default:
		fmt.Println("Invalid input, choosing default option of automaticaly selecting:" + defaultPlaylist)
	}

	return &Config{
		General: General{
			User:            user,
			DefaultPlaylist: defaultPlaylist,
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

// GenerateDefaultConfig creates and saves the default configuration file.
// Returns true if successful, false otherwise.
func GenerateDefaultConfig() bool {
	configPath := filepath.Join(dirConfig, fileConfig)

	if !SaveConfig(configPath, DefaultConfig()) {
		logger.Log.Error("Generate default configuration failed")
		return false
	}
	logger.Log.Info("Generate default configuration success")
	return true
}
