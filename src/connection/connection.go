package connection

import (
	"log"

	"Melodex/src/backend/music"
	"Melodex/src/backend/playlist"

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
	rows, err := music.Songlist(playlistPath)
	if err != nil {
		log.Println("Error connecting songs:", err)
		return nil
	}
	allSongs = rows
	return rows
}

func Play() {
	if len(allSongs) == 0 {
		log.Println("No songs available to play.")
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
		log.Println("Invalid mode specified.")
		return
	}

	music.PlaySong(song)
}
