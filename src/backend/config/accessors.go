package config

import (
	"Melodex/src/log"
	"os"
	"strconv"
	"strings"
)

func GetVolume() int {
	volume, err := strconv.Atoi(strings.TrimSuffix(loadedConfig.General.DefaultVolume, "%"))
	if err != nil || volume < 0 || volume > 100 {
		log.Log.Warn("Invalid volume format in config; using default (100%)")
		return 100
	}
	return volume
}

func GetUser() string {
	if loadedConfig.General.User != "" {
		return loadedConfig.General.User
	}
	user := os.Getenv("USER")
	if user == "" {
		user = "unknown"
	}
	log.Log.Infof("User set to: %s", user)
	return user
}

func GetDefaultPlaylist() string {
	if loadedConfig.General.DefaultPlaylist == "" {
		log.Log.Warn("No default playlist found; returning empty string")
		return ""
	}
	return loadedConfig.General.DefaultPlaylist
}

func GetPlaylistArray() []string {
	if len(loadedConfig.Playlist.Playlists) == 0 {
		log.Log.Warn("No playlists available in config")
	}
	return loadedConfig.Playlist.Playlists
}

func GetBorderColor() string {
	if loadedConfig.Design.Border == "" {
		log.Log.Warn("No border color defined; using default (#FFFFFF)")
		return "#FFFFFF"
	}
	return loadedConfig.Design.Border
}

func GetForegroundColor() string {
	if loadedConfig.Design.Foreground == "" {
		log.Log.Warn("No foreground color defined; using default (#FFFFFF)")
		return "#FFFFFF"
	}
	return loadedConfig.Design.Foreground
}

func GetBackgroundColor() string {
	if loadedConfig.Design.Background == "" {
		log.Log.Warn("No background color defined; using default (#000000)")
		return "#000000"
	}
	return loadedConfig.Design.Background
}
