package cache

import (
	"Melodex/src/log"
	"Melodex/src/settings"
	"encoding/json"
	"github.com/charmbracelet/bubbles/table"
	"os"
)

func SaveCache(playlist string, content []table.Row) error {
	cachePath := settings.AppConfig.CacheDir + playlist + ".cache"
	data, err := json.Marshal(content)
	if err != nil {
		log.Log.Errorf("Error writing cache file: %v", err)
		return err
	}
	return os.WriteFile(cachePath, data, 0644)
}
