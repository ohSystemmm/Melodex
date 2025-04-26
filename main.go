package main

import (
	"fmt"
	"os"

	"Melodex/src/backend/music"
	tui "Melodex/src/tui"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: melodex <command>")
		return
	}

	greeter()

	command := os.Args[1]
	switch command {
	case "init", "-i":
		fmt.Println("Initializing Melodex...")
	case "--version", "-v":
		fmt.Println("Melodex version 0.1.0")
		return
	case "--help", "-h":
		help()
		return
	case "--playlist", "-p":
		fmt.Println("Playlist option selected")
	case "--config", "-c":
		fmt.Println("Config option selected")
	case "--debug", "-d":
		fmt.Println("Debug option selected")
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Use --help for a list of available commands.")
		return
	}

	music.Init()
	tui.Application()
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
