package music

import (
	"Melodex/src/logger"
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// FUNCTION: Gets cache path
func getCachePath(cacheDir, playlistName string) string {
	return filepath.Join(cacheDir, playlistName+".cache")
}

// FUNCTION: Loads cached data
func loadCache(cacheDir, playlistName string) ([]table.Row, error) {
	path := getCachePath(cacheDir, playlistName)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var rows []table.Row
	if err := json.Unmarshal(data, &rows); err != nil {
		return nil, err
	}

	return rows, nil
}

// FUNCTION: Saves cache
func saveCache(cacheDir, playlistName string, rows []table.Row) error {
	path := getCachePath(cacheDir, playlistName)
	data, err := json.Marshal(rows)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// FUNCTION: Gets duration of a song
func getDuration(songPath string) (string, error) {
	output, err := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration",
		"-of", "default=noprint_wrappers=1:nokey=1", songPath).Output()
	if err != nil {
		return "", err
	}

	seconds, err := strconv.ParseFloat(strings.TrimSpace(string(output)), 64)
	if err != nil {
		return "", err
	}

	minutes := int(seconds) / 60
	seconds = math.Mod(seconds, 60)

	return fmt.Sprintf("%02d:%02d", minutes, int(seconds)), nil
}

// FUNCTION: Gets all song names in a directory
func getAllSongNames(directory string) ([]string, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var songNames []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		songNames = append(songNames, file.Name())
	}

	return songNames, nil
}

// FUNCTION: Gets durations of all songs in a directory
func getAllDurations(directory string) ([]string, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var durations []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		duration, err := getDuration(filepath.Join(directory, file.Name()))
		if err != nil {
			logger.Log.Warnf("Error fetching duration for %s: %v", file.Name(), err)
			durations = append(durations, "Unknown")
		} else {
			durations = append(durations, duration)
		}
	}

	return durations, nil
}

// FUNCTION: Generates song list and applies caching
func GenerateSongList(directory, cacheDir string) ([]table.Row, error) {
	playlistName := filepath.Base(directory)
	rows, err := loadCache(cacheDir, playlistName)
	if err == nil {
		return rows, nil
	}

	songNames, err := getAllSongNames(directory)
	if err != nil {
		return nil, err
	}

	durations, err := getAllDurations(directory)
	if err != nil {
		return nil, err
	}

	var songList []table.Row
	for i := range songNames {
		songList = append(songList, table.Row{songNames[i], durations[i]})
	}

	// Cache the results
	if err := saveCache(cacheDir, playlistName, songList); err != nil {
		logger.Log.Warnf("Warning: Could not save playlist cache: %v", err)
	}

	return songList, nil
}
