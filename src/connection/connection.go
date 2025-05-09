package connection

import (
	"Melodex/src/backend/music"
	"log"

	"github.com/charmbracelet/bubbles/table"
)

var (
	playlistPath string
)

// Services but name is cooler

func GetPlaylistPath(param string) {
	playlistPath = param
}

func ConnectSongs() []table.Row {

	if playlistPath == "" {
		// playlistPath = music.GetDefaultPlaylistPath()
		return nil
	}

	rows, err := music.Songlist(playlistPath)
	if err != nil {
		log.Println("Error connecting songs:", err)
		return nil
	}
	return rows
}
