package connection

import (
	"Melodex/src/backend/config"
	"Melodex/src/backend/music"
	"Melodex/src/logger"
	"github.com/charmbracelet/bubbles/table"
)

func ConnectSongs() []table.Row {
	var err error
	songList, err = music.GenerateSongList(playlistPath, "/home/"+config.ConfGetUser()+"/.cache/melodex")
	if err != nil {
		logger.Log.Errorf("Error connecting to playlist: %v", err)
	}

	return songList
}

func Play() {
	music.PlaySong(playlistPath + "/" + currentSong)
}
