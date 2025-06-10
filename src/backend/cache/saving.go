package cache

import (
	"Melodex/src/log"
	"Melodex/src/settings"
	"encoding/json"
	"github.com/charmbracelet/bubbles/table"
	"os"
	"path/filepath"
)

func SaveCache(playlist string, content []table.Row) error {
	cachePath := settings.AppConfig.CacheDir + playlist + ".cache"
	cacheDir := filepath.Dir(cachePath) // Extract the directory path

	// Ensure the cache directory exists
	if err := os.MkdirAll(cacheDir, os.ModePerm); err != nil {
		log.Log.Errorf("Error creating cache directory: %v", err)
		return err
	}

	// Check if the cache file exists
	if _, err := os.Stat(cachePath); err == nil {
		// Cache file exists, check if it's valid
		if isCacheValid(cachePath) {
			log.Log.Infof("Cache file already exists and is valid: %s", cachePath)
			return nil // No need to recreate
		}
		log.Log.Warnf("Cache file is corrupt, regenerating: %s", cachePath)
	} else if !os.IsNotExist(err) {
		log.Log.Errorf("Error checking cache file: %v", err)
		return err
	}

	// Convert content to JSON
	data, err := json.Marshal(content)
	if err != nil {
		log.Log.Errorf("Error marshalling cache content: %v", err)
		return err
	}

	// Write the cache file
	if err := os.WriteFile(cachePath, data, 0644); err != nil {
		log.Log.Errorf("Error writing cache file: %v", err)
		return err
	}

	log.Log.Infof("Cache file created: %s", cachePath)
	return nil
}

// Function to check if cache file is valid
func isCacheValid(cachePath string) bool {
	data, err := os.ReadFile(cachePath)
	if err != nil {
		log.Log.Errorf("Error reading cache file: %v", err)
		return false
	}

	var content []table.Row
	if err := json.Unmarshal(data, &content); err != nil {
		log.Log.Errorf("Cache file is corrupt: %v", err)
		return false
	}

	return true // File is valid
}
