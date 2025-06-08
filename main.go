package main

import (
	"fmt"
	"os"
	"path/filepath"

	"Melodex/src/backend/config"
	"Melodex/src/backend/music"
	"Melodex/src/connection"
	"Melodex/src/logger"
	"Melodex/src/tui"
)

var (
	version  = "0.0.8"                                              // Defines the current application version.
	release  = "2025-XX-XX"                                         // Placeholder for the release date.
	cacheDir = "/home/" + config.ConfGetUser() + "/.cache/melodex/" // Defines the cache directory path.
)

// main is the entry point for Melodex.
func main() {
	displayFileContent("src/assets/usBTW.txt", "Welcome to Melodex!")

	if len(os.Args) < 2 {
		startApplication()
		return
	}

	command := os.Args[1]
	handleCommand(command)
}

// startApplication initializes the application, checks config files, and starts the UI.
func startApplication() {
	configPath := "/home/" + config.ConfGetUser() + "/.config/melodex/config.toml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		logger.Log.Info("Config file not found. Creating defaults.")
		config.GenerateDefaultConfig()
	} else {
		cfg := config.LoadConfig(configPath)
		if err != nil {
			logger.Log.Info("Config file corrupted. Creating defaults.")
			config.GenerateDefaultConfig()
		} else {
			logger.Log.Info("Config file found at " + configPath)
			_ = cfg
		}
	}

	connection.SetPlaylistPath(config.ConfGetDefaultPlaylist())
	music.Init()
	tui.Application()
}

// handleCommand processes command-line arguments and executes corresponding actions.
func handleCommand(command string) {
	switch command {
	case "--version", "-v":
		fmt.Printf("Melodex version %s\nRelease date: %s\nUse --help for a list of available commands.\n", version, release)
	case "--help", "-h":
		displayFileContent("src/assets/help.txt", "Options not found, check out https://github.com/ohSystemmm/Melodex")
	case "--playlist", "-p":
		handlePlaylist()
		startApplication()
	case "--clear-cache", "-cc":
		err := removeCache()
		if err != nil {
			logger.Log.Warn("Failed to remove cache: " + err.Error())
		}
		fmt.Println("Removed cache")
	case "--config", "-c":
		handleConfig()
		startApplication()
	case "--default-config", "-dc":
		regenerateDefaultConfig()
		startApplication()
	default:
		fmt.Printf("Unknown command: %s\nUse --help for a list of available commands.\n", command)
	}
}

// handlePlaylist sets the playlist path based on command-line input.
func handlePlaylist() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: melodex --playlist <path>")
		return
	}
	playlistPath := os.Args[2]
	connection.SetPlaylistPath(playlistPath)
}

// handleConfig loads the specified configuration file from command-line input.
func handleConfig() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: melodex --config <path>")
		return
	}
	configPath := os.Args[2]
	fmt.Printf("Loading config: %s\n", configPath)
	config.LoadConfig(configPath)
}

// regenerateDefaultConfig resets the application to its default configuration.
func regenerateDefaultConfig() {
	logger.Log.Info("Regenerating default configuration...")
	if config.GenerateDefaultConfig() {
		logger.Log.Info("All configurations successfully regenerated.")
	} else {
		logger.Log.Error("Failed to regenerate the default configuration.")
	}
	startApplication()
}

// removeCache deletes cached files in the application's cache directory.
func removeCache() error {
	logger.Log.Info("Removing cache...")

	files, err := os.ReadDir(cacheDir)
	if err != nil {
		logger.Log.Errorf("Error reading cache directory: %v", err)
		return err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".cache" {
			filePath := filepath.Join(cacheDir, file.Name())
			err := os.Remove(filePath)
			if err != nil {
				logger.Log.Warnf("Error removing cache file %s: %v", file.Name(), err)
			} else {
				logger.Log.Infof("Removed cache file: %s", file.Name())
			}
		}
	}
	return nil
}

// displayFileContent reads and prints file content or an error message if reading fails.
func displayFileContent(path string, errorMessage string) {
	file, err := os.ReadFile(path)
	if err != nil {
		logger.Log.Warningf("Error reading file %s: %v\n", path, err)
		fmt.Println(errorMessage)
		return
	}
	fmt.Println(string(file))
}
