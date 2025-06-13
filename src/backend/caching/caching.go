package caching

import (
	"Melodex/src/settings"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"Melodex/src/log"
)

func getDuration(songPath string) string {
	output, err := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration",
		"-of", "default=noprint_wrappers=1:nokey=1", songPath).Output()
	if err != nil {
		log.Log.Errorf("Error getting duration: %v", err)
		return "Unknown"
	}

	seconds, err := strconv.ParseFloat(strings.TrimSpace(string(output)), 64)
	if err != nil {
		log.Log.Errorf("Error parsing duration: %v", err)
		return "Unknown"
	}

	minutes := int(seconds) / 60
	seconds = math.Mod(seconds, 60)

	return fmt.Sprintf("%02d:%02d", minutes, int(seconds))
}

func getAllSongNames() ([]string, error) {
	files, err := os.ReadDir(settings.AppConfig.PlaylistPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var songNames []string
	for _, file := range files {
		if !file.IsDir() {
			songNames = append(songNames, file.Name())
		}
	}

	return songNames, nil
}

func getAllDurations() ([]string, error) {
	files, err := os.ReadDir(settings.AppConfig.PlaylistPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var durations []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		durations = append(durations, getDuration(filepath.Join(settings.AppConfig.PlaylistPath, file.Name())))
	}

	return durations, nil
}
