package cache

import (
	"Melodex/src/backend/config"
	"Melodex/src/log"
	"Melodex/src/settings"
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	"os"
)

var index int

func LoadStoredPlaylist(passedIndex int) {
	index = passedIndex
}

func loadConfiguredCache() ([]table.Row, error) {
	playlists := config.GetPlaylists()

	if index < 0 || index >= len(playlists) {
		log.Log.Warnf("Playlist index out of range: %d", index)
		return nil, fmt.Errorf("invalid playlist index: %d", index)
	}

	return LoadCache(playlists[index])
}

func LoadCache(playlist string) ([]table.Row, error) {
	data, err := os.ReadFile(settings.AppConfig.CacheDir + playlist + ".cache")
	if err != nil {
		log.Log.Errorf("Error reading cache file: %v", err)
		return nil, err
	}

	var rows []table.Row
	if err = json.Unmarshal(data, &rows); err != nil {
		return nil, err
	}
	return rows, nil
}
