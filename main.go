package main

import (
	"fmt"
	"log"
	"os"

	"Melodex/src/backend/config"
	"Melodex/src/backend/music"
	"Melodex/src/connection"
	tui "Melodex/src/tui"
)

var (
	version = "0.0.8"
	release = "2025-XX-XX"
)

func main() {
	if len(os.Args) < 2 {
		startApplication()
		return
	}

	command := os.Args[1]
	handleCommand(command)
}
func startApplication() {
	configPath := os.Getenv("HOME") + "/.config/melodex/config.toml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("Configuration file not found. Generating default configuration...")
		config.GenerateDefaultConfig()
	} else {
		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			fmt.Println("Configuration file is corrupted. Generating default configuration...")
			config.GenerateDefaultConfig()
		} else {
			fmt.Printf("Loading configuration from %s...\n", configPath)
			_ = cfg
		}
	}

	music.Init()
	tui.Application()
}

func handleCommand(command string) {
	switch command {
	case "init", "i":
		fmt.Println("Initializing Melodex...")
		startApplication()
	case "--version", "-v":
		displayVersion()
	case "--help", "-h":
		displayHelp()
	case "--playlist", "-p":
		handlePlaylist()
		startApplication()
	case "--config", "-c":
		handleConfig()
		startApplication()
	case "--debug", "-d":
		startDebugMode()
	case "--default-config", "-dc":
		regenerateDefaultConfig()
		startApplication()
	default:
		fmt.Printf("Unknown command: %s\nUse --help for a list of available commands.\n", command)
	}
}

func displayVersion() {
	fmt.Printf("Melodex version %s\nRelease date: %s\nUse --help for a list of available commands.\n", version, release)
}

func displayHelp() {
	displayFileContent("src/assets/help.txt", "Help file not found")
}

func handlePlaylist() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: melodex --playlist <path>")
		return
	}
	playlistPath := os.Args[2]
	connection.SetPlaylistPath(playlistPath)
}

func handleConfig() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: melodex --config <path>")
		return
	}
	configPath := os.Args[2]
	fmt.Printf("Loading config: %s\n", configPath)
	config.LoadConfig(configPath)
}

func startDebugMode() {
	fmt.Println("Starting Melodex in debug mode...")
	// TODO: Implement debug mode (logs to stdout instead of log file)
}

func regenerateDefaultConfig() {
	fmt.Println("Regenerating default configuration...")
	if config.GenerateDefaultConfig() {
		fmt.Println("Default configuration successfully regenerated.")
	} else {
		fmt.Println("Failed to regenerate default configuration.")
	}
	startApplication()
}

func displayFileContent(path string, errorMessage string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Printf("%s: %v\n", errorMessage, err)
		return
	}
	fmt.Println(string(file))
}
