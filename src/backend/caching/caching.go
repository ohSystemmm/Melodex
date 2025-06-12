package caching

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"Melodex/src/log"
)

func getCachePath(cacheDir, playlistName string) string {
	return filepath.Join(cacheDir, playlistName+".cache")
}

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

func getAllSongNames(directory string) ([]string, error) {
	files, err := os.ReadDir(directory)
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

func getAllDurations(directory string) ([]string, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var durations []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		durations = append(durations, getDuration(filepath.Join(directory, file.Name())))
	}

	return durations, nil
}
