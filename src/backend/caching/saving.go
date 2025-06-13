package caching

import (
	"Melodex/src/settings"
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	"os"
)

func saveCache(rows []table.Row) error {
	if err := os.MkdirAll(settings.AppConfig.CacheDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}

	data, err := json.Marshal(rows)
	if err != nil {
		return fmt.Errorf("failed to serialize cache data: %w", err)
	}

	return os.WriteFile(settings.AppConfig.CacheDir+settings.AppConfig.PlaylistName+".cache", data, 0644)
}
