package config

import (
	"Melodex/src/backend/settings"
	"Melodex/src/log"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func DefaultConfig() *Config {
	user := os.Getenv("USER")
	if user == "" {
		user = "unknown"
	}

	basePath := filepath.Join("/home", user)
	defaultPlaylist := filepath.Join(basePath, "example_playlist")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Set default playlist? (y/n): ")
	input, _ := reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))

	if input == "y" || input == "yes" {
		fmt.Print("Enter playlist path (e.g., /home/user/music): ")
		customPlaylist, _ := reader.ReadString('\n')
		defaultPlaylist = strings.TrimSpace(customPlaylist)
	}

	return &Config{
		General: General{
			User:            user,
			DefaultPlaylist: defaultPlaylist,
			DefaultVolume:   "100%",
		},
		Playlist: Playlist{
			Playlists: []string{
				filepath.Join(basePath, "playlist1"),
				filepath.Join(basePath, "playlist2"),
				filepath.Join(basePath, "playlist3"),
			},
		},
		Design: Design{
			Border:     "#FFFFFF",
			Foreground: "#FFFFFF",
			Background: "#000000",
		},
	}
}

func GenerateDefaultConfig() error {
	configPath := settings.AppConfig.ConfigDir + settings.AppConfig.ConfigFile
	if err := SaveConfig(configPath, DefaultConfig()); err != nil {
		log.Log.Errorf("Failed to generate default config: %v", err)
		return err
	}

	log.Log.Info("Default config generated successfully")
	return nil
}
