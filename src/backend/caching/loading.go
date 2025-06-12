package caching

import (
	"Melodex/src/log"
	"encoding/json"
	"github.com/charmbracelet/bubbles/table"
	"os"
)

func loadCache(cachePath string) ([]table.Row, error) {
	data, err := os.ReadFile(cachePath)
	if err != nil {
		log.Log.Errorf("Error loading cache file: %v", err)
		return nil, err
	}

	var rows []table.Row
	if err := json.Unmarshal(data, &rows); err != nil {
		log.Log.Errorf("Error parsing cache file: %v", err)
		return nil, err
	}

	return rows, nil
}
