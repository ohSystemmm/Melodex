package main

import (
	"Melodex/src/settings"
	"fmt"
	"os"
	"path/filepath"

	"Melodex/src/backend/config"
	"Melodex/src/backend/music"
	"Melodex/src/log"
	"Melodex/src/tui"
)

func main() {
	displayFileContent("src/assets/melodex.txt", "Welcome to Melodex!")
	config.InitConfig()

	if len(os.Args) < 2 {
		startApplication()
		return
	}

	command := os.Args[1]
	handleCommand(command)
}

func startApplication() {
	cfgPath := "/home/ohsystemmm/.config/melodex/melodex.toml"

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Log.Info("Config file not found. Creating defaults.")

		err = config.GenerateDefaultConfig()
		if err != nil {
			log.Log.Errorf("Error creating default configuration: %v", err)
		}
	} else {
		cfg := config.LoadConfig()
		if cfg == nil {
			log.Log.Warn("Config file corrupted. Creating defaults.")
			err = config.GenerateDefaultConfig()
			if err != nil {
				log.Log.Errorf("Error creating default configuration: %v", err)
			}
		} else {
			log.Log.Info("Config file found at " + cfgPath)
		}
	}
	music.InitVLC()
	tui.Application()
}

func handleCommand(command string) {
	switch command {
	case "--version", "-v":
		fmt.Printf("Melodex version %s\nRelease date: %s\nUse --help for a list of available commands.\n",
			settings.AppConfig.Version, settings.AppConfig.Release)

	case "--help", "-h":
		displayFileContent("src/assets/help.txt",
			"Options not found, check out https://github.com/ohSystemmm/Melodex")

	case "--playlist", "-p":
		handlePlaylist()

	case "--clear-cache", "-cc":
		if err := removeCache(); err != nil {
			log.Log.Warnf("Failed to remove cache: %v", err)
		} else {
			fmt.Println("Cache removed successfully")
		}

	case "--config", "-c":
		handleConfig()

	case "--default-config", "-dc":
		regenerateDefaultConfig()

	default:
		fmt.Printf("Unknown command: %s\nUse --help for a list of available commands.\n", command)
	}
}

func handlePlaylist() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: melodex --playlist <path>")
		return
	}
	playlistPath := os.Args[2]

	settings.SetPlaylistPath(playlistPath)
	settings.SetPlaylistName(filepath.Base(playlistPath))

	startApplication()
}

func handleConfig() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: melodex --config <path>")
		return
	}
	configPath := os.Args[2]
	settings.SetConfigFile(configPath)
	config.LoadConfig()
	fmt.Printf("Config Loaded: %s\n", configPath)
	startApplication()
}

func regenerateDefaultConfig() {
	log.Log.Info("Regenerating default config...")
	if err := config.GenerateDefaultConfig(); err != nil {
		log.Log.Errorf("Failed to regenerate default config: %v", err)
	}
	startApplication()
}

func removeCache() error {
	log.Log.Info("Removing cache...")
	err := os.RemoveAll(settings.AppConfig.CacheDir)
	if err != nil {
		return fmt.Errorf("error removing cache directory: %w", err)
	}
	return nil
}

func displayFileContent(path string, errorMessage string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Log.Warnf("Error reading file %s: %v", path, err)
		fmt.Println(errorMessage)
		return
	}
	fmt.Println(string(file))
}
