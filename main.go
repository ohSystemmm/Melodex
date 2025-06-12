package main

import (
	"Melodex/src/backend/config"
	"Melodex/src/backend/music"
	"Melodex/src/log"
	"Melodex/src/settings"
	"Melodex/src/tui"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Print(printFileContent("src/assets/melodex.txt", "Welcome to Melodex!"))
	if len(os.Args) < 2 {
		startMelodex()
		return
	}

	handleCmd(os.Args[1])
}

func startMelodex() {
	music.InitVLC()
	tui.Application()
}

func handleCmd(cmd string) {
	switch cmd {
	case "init":
		settings.SetupMelodex()
	case "-h", "--help":
		fmt.Print(printFileContent("src/assets/help.txt",
			"Options not found, check out https://github.com/ohSystemmm/Melodex"))
	case "-v", "--version":
		fmt.Printf("Melodex version %s\nRelease date: %s\n\nUse --help for a list of available commands.\n\n",
			settings.AppConfig.Version, settings.AppConfig.Release)
	case "-c", "--config":
		path := parsePath()
		config.LoadConfig(path)
		//config.LoadConfig(path)
	case "-p", "--playlist":
		//path := parsePath()
		//playlist.UsePlaylist(path)
	case "-cc", "--clear-cache":
		//cache.DelCache()
	case "-dc", "--default-config":
		config.SaveConfig(config.GenerateDefaultConfig())
	case "-sp", "--select-playlist":
		//if len(os.Args) < 4 {
		//	log.Log.Error("Missing playlist index argument.")
		//	return
		//}
		//index, err := strconv.Atoi(os.Args[3])
		//if err != nil {
		//	log.Log.Error("Invalid playlist index:", err)
		//	return
		//}
		//cache.LoadStoredPlaylist(index)
	default:
		fmt.Printf("Unknown command: %s\n\nUse --help for a list of available commands.\n\n", cmd)
	}
}

func parsePath() string {
	if len(os.Args) < 3 {
		fmt.Printf("Error: No path provided\n\nUse --help for a list of available commands.\n\n")
		return ""
	}
	path := os.Args[2]
	if path[:2] == "~/" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Log.Errorf("Error getting user home directory: %v", err)
			return ""
		}
		path = filepath.Join(homeDir, path[2:])
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Log.Errorf("Error converting to absolute path: %v", err)
		return ""
	}
	return absPath
}

func printFileContent(file string, errorMessage string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Log.Errorf("Error reading file %s: %v", file, err)
		fmt.Println(errorMessage)
		return ""
	}
	return string(content)
}
