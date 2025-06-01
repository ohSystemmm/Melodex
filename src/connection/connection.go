package connection

import (
	"Melodex/src/backend/music"
	"Melodex/src/backend/playlist"
	"Melodex/src/logger"

	"github.com/charmbracelet/bubbles/table"
)

var (
	playlistPath string
	currentSong  string
	shuffle      bool
	mode         int
	allSongs     []table.Row
)

func SetPlaylistPath(passedPlaylistPath string) {
	playlistPath = passedPlaylistPath
}

func SetCurrentSong(song string) {
	currentSong = song
}

func GetCurrentSong() string {
	return currentSong
}

func ConnectSongs() []table.Row {
	allSongs = music.SongList(playlistPath)
	return allSongs
}

func Play() {
	if len(allSongs) == 0 {
		logger.Log.Warn("No songs found")
		return
	}

	playlist.GetAllSongs(allSongs)

	var song string
	switch {
	case mode < 0:
		song = playlistPath + "/" + currentSong
	case mode == 0:
		if shuffle {
			song = playlistPath + "/" + playlist.GetRandomSong()
		} else {
			song = playlistPath + "/" + playlist.GetNextSong()
		}
	default:
		logger.Log.Error("Invalid Mode")
		return
	}

	music.PlaySong(song)
}
