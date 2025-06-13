package caching

import (
	"Melodex/src/log"
	"github.com/charmbracelet/bubbles/table"
)

func GenerateSongList() ([]table.Row, error) {
	rows, err := loadCache()
	if err == nil {
		log.Log.Info("Cache loaded")
		return rows, nil
	}

	songNames, err := getAllSongNames()
	if err != nil {
		return nil, err
	}

	durations, err := getAllDurations()
	if err != nil {
		return nil, err
	}

	songList := make([]table.Row, len(songNames))
	for i := range songNames {
		songList[i] = table.Row{songNames[i], durations[i]}
	}

	if err = saveCache(songList); err != nil {
		log.Log.Warnf("Warning: Could not save playlist cache: %v", err)
	}

	return songList, nil
}
