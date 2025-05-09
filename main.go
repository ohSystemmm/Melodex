package main

import (
	"fmt"
	"os"
	"strings"

	"Melodex/src/backend/music"
	"Melodex/src/connection"
	tui "Melodex/src/tui"
)

var (
	version = "0.3.1"
	release = "2025-XX-XX"
)

func main() {
	printGreeter()

	if len(os.Args) < 2 {
		startApp()
		return
	}

	command := os.Args[1]
	switch command {
	case "init", "i":
		// TODO: Implement initialization logic
		startApp()
	case "--version", "-v":
		fmt.Println("Melodex version", version)
		fmt.Println("Release date:", release)
		fmt.Println("Use --help for a list of available commands.")
	case "--help", "-h":
		printFileContent("src/assets/help.txt")
	case "--playlist", "-p":
		if len(os.Args) > 2 {
			path := handlePaths(os.Args[2])
			if path == "" {
				fmt.Println("Error: Invalid path provided.")
				os.Exit(1)
			}
			connection.GetPlaylistPath(path)
			startApp()
		} else {
			fmt.Println("Error: Missing playlist path after -p or --playlist")
			os.Exit(1)
		}
	case "--config", "-c":
		if len(os.Args) > 2 {
			path := handlePaths(os.Args[2])
			if path == "" {
				fmt.Println("Error: Invalid path provided.")
				os.Exit(1)
			}
			connection.GetPlaylistPath(path)
		}
		startApp()
	case "--debug", "-d":
		// TODO: Opens melodex in debug mode
		startApp()
	case "--default-config", "-dc":
		// TODO: Regenerate the default config and opens it
		startApp()
	default:
		fmt.Fprintln(os.Stderr, "Unknown command: ", command)
		fmt.Println("Use --help for a list of available commands.")
		os.Exit(1)
	}
}

func handlePaths(argument string) string {
	argument = strings.TrimSpace(argument) // Trim leading/trailing spaces
	if strings.HasPrefix(argument, "\"") && strings.HasSuffix(argument, "\"") && len(argument) > 1 {
		argument = argument[1 : len(argument)-1]
	} else if strings.HasPrefix(argument, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error getting home directory:", err)
			os.Exit(1)
		}
		argument = strings.Replace(argument, "~", homeDir, 1)
	}
	return argument
}

func startApp() {
	music.Init()
	tui.Application()
}

func printGreeter() {
	printFileContent("src/assets/usBTW.txt")
}

func printFileContent(filePath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
		return // Don't os.Exit here, just print error
	}
	fmt.Println("\n\033[36m" + string(content) + "\033[0m\n")
}
