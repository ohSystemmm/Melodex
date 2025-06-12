package caching

import (
	"Melodex/src/log"
	"github.com/charmbracelet/bubbles/table"
	"path/filepath"
)

func GenerateSongList(directory, cacheDir string) ([]table.Row, error) {
	playlistName := filepath.Base(directory)
	cachePath := getCachePath(cacheDir, playlistName)

	rows, err := loadCache(cachePath)
	if err == nil {
		log.Log.Infof("Cache loaded: %s", cachePath)
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

	songList := make([]table.Row, len(songNames))
	for i := range songNames {
		songList[i] = table.Row{songNames[i], durations[i]}
	}

	if err = saveCache(cachePath, songList); err != nil {
		log.Log.Warnf("Warning: Could not save playlist cache: %v", err)
	}

	return songList, nil
}
