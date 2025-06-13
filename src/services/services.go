package services

import (
	"Melodex/src/backend/caching"
	"Melodex/src/log"
	"Melodex/src/settings"
	"github.com/charmbracelet/bubbles/table"
)

func GenerateMusicList() []table.Row {
	songList, err := caching.GenerateSongList()
	if err != nil {
		log.Log.Errorf("Error connecting to playlist: %v", err)
	}

	settings.SetSongList(songList)
	return songList
}
