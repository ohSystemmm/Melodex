package cache

import (
	"Melodex/src/log"
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

func getCachePath(cacheDir, playlistName string) string {
	return filepath.Join(cacheDir, playlistName+".cache")
}

func loadCache(cacheDir, playlistName string) ([]table.Row, error) {
	path := getCachePath(cacheDir, playlistName)

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var rows []table.Row
	if err := json.Unmarshal(data, &rows); err != nil {
		log.Log.Warnf("Cache file is corrupt, regenerating: %s", path)
		return nil, err
	}

	return rows, nil
}

func saveCache(cacheDir, playlistName string, rows []table.Row) error {
	path := getCachePath(cacheDir, playlistName)

	// Ensure cache directory exists
	if err := os.MkdirAll(cacheDir, os.ModePerm); err != nil {
		log.Log.Errorf("Error creating cache directory: %v", err)
		return err
	}

	data, err := json.Marshal(rows)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

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

	hours := int(seconds) / 3600
	minutes := (int(seconds) / 60) % 60
	seconds = math.Mod(seconds, 60)

	// Format based on length
	if hours > 0 {
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, int(seconds)), nil
	}
	return fmt.Sprintf("%02d:%02d", minutes, int(seconds)), nil
}

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
		songNames = append(songNames, filepath.Join(directory, file.Name()))
	}

	return songNames, nil
}

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
			log.Log.Warnf("Error fetching duration for %s: %v", file.Name(), err)
			durations = append(durations, "Unknown")
		} else {
			durations = append(durations, duration)
		}
	}

	return durations, nil
}

func GenerateSongList(directory, cacheDir string) ([]table.Row, error) {
	playlistName := filepath.Base(directory)
	rows, err := loadCache(cacheDir, playlistName)
	if err == nil {
		log.Log.Infof("Cache file loaded successfully: %s", getCachePath(cacheDir, playlistName))
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

	if err = saveCache(cacheDir, playlistName, songList); err != nil {
		log.Log.Warnf("Warning: Could not save playlist cache: %v", err)
	}

	return songList, nil
}
