package caching

import (
	"Melodex/src/log"
	"Melodex/src/settings"
	"encoding/json"
	"github.com/charmbracelet/bubbles/table"
	"os"
)

func loadCache() ([]table.Row, error) {
	data, err := os.ReadFile("/home/ohsystemmm/.cache/melodex/cache/" + settings.AppConfig.PlaylistName + ".cache")
	if err != nil {
		log.Log.Errorf("Error loading cache file: %v", err)
		return nil, err
	}

	var rows []table.Row
	if err = json.Unmarshal(data, &rows); err != nil {
		log.Log.Errorf("Error parsing cache file: %v", err)
		return nil, err
	}

	return rows, nil
}
