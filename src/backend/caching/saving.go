package caching

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	"os"
	"path/filepath"
)

func saveCache(cachePath string, rows []table.Row) error {
	cacheDir := filepath.Dir(cachePath)
	if err := os.MkdirAll(cacheDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}

	data, err := json.Marshal(rows)
	if err != nil {
		return fmt.Errorf("failed to serialize cache data: %w", err)
	}

	return os.WriteFile(cachePath, data, 0644)
}
