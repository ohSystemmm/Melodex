package main

import (
	"fmt"
	"os"

	"Melodex/src/backend/music"
	tui "Melodex/src/tui"
)

var (
	version = "0.0.8"
	release = "2025-XX-XX"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: melodex <command>")
		tui.Application()
		music.Init()
		return
	}

	greeter()

	command := os.Args[1]
	switch command {
	case "init", "i":
	case "--version", "-v":
		fmt.Println("Melodex version", version)
		fmt.Println("Release date:", release)
		fmt.Println("Use --help for a list of available commands.")
		return
	case "--help", "-h":
		help()
		return
	case "--playlist", "-p":
		// TODO: Opens melodex with the specified playlist
	case "--config", "-c":
		// TODO: Opens melodex with the specified config
	case "--debug", "-d":
		// TODO: Opens melodex in debug mode
	case "--default-config", "-dc":
		// TODO: Regenerate the default config and opens it

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Use --help for a list of available commands.")
		return
	}

	greeter()
}

func greeter() {
	file, err := os.ReadFile("src/assets/usBTW.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n\033[36m" + string(file) + "\033[0m\n")
}

func help() {
	file, err := os.ReadFile("src/assets/help.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file) + "\n")
}
