package cache

import (
	"Melodex/src/log"
	"Melodex/src/settings"
	"os"
)

func DelCache() {
	err := os.RemoveAll(settings.AppConfig.CacheDir)
	if err != nil {
		log.Log.Errorf("Failed to delete cache folder '%s': %v\n", settings.AppConfig.CacheDir, err)
	}
}

func GetDurations() []string {
	return nil
}
