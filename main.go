package main

import (
	"fmt"
	"os"
	"path/filepath"

	"Melodex/src/backend/config"
	"Melodex/src/backend/music"
	"Melodex/src/backend/settings"
	"Melodex/src/log"
	"Melodex/src/services"
	"Melodex/src/tui"
)

func main() {
	displayFileContent("src/assets/usBTW.txt", "Welcome to Melodex!")

	if len(os.Args) < 2 {
		startApplication()
		return
	}

	command := os.Args[1]
	handleCommand(command)
}

func startApplication() {
	configPath := filepath.Join(settings.AppConfig.ConfigDir, settings.AppConfig.ConfigFile)

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Log.Info("Config file not found. Creating defaults.")
		if err := config.GenerateDefaultConfig(); err != nil {
			log.Log.Errorf("Error creating default configuration: %v", err)
		}
	} else {
		cfg := config.LoadConfig()
		if cfg == nil {
			log.Log.Warn("Config file corrupted. Creating defaults.")
			if err := config.GenerateDefaultConfig(); err != nil {
				log.Log.Errorf("Error creating default configuration: %v", err)
			}
		} else {
			log.Log.Info("Config file found at " + configPath)
		}
	}

	services.SetPlaylistPath(config.GetDefaultPlaylist())
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
	services.SetPlaylistPath(playlistPath)
	startApplication()
}

func handleConfig() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: melodex --config <path>")
		return
	}
	configPath := os.Args[2]
	fmt.Printf("Loading config: %s\n", configPath)
	settings.SetConfigFile(configPath)
	config.LoadConfig()
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

	files, err := os.ReadDir(settings.AppConfig.CacheDir)
	if err != nil {
		return fmt.Errorf("error reading cache directory: %w", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".cache" {
			filePath := filepath.Join(settings.AppConfig.CacheDir, file.Name())
			if err := os.Remove(filePath); err != nil {
				log.Log.Warnf("Error removing cache file %s: %v", file.Name(), err)
			} else {
				log.Log.Infof("Removed cache file: %s", file.Name())
			}
		}
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
