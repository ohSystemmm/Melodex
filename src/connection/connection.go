package connection

import (
	"Melodex/src/backend/music"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/table"
)

// Services but name is cooler

// var tempPlaylist = "/home/" + os.Getenv("USER") + "/Music/"
var tempPlaylist = "/home/" + os.Getenv("USER") + "/HolyMoly/008_Music/Best_Songs_Ever-ohSystemmm/"

func GetPlaylistPath() string {
	return tempPlaylist
}

func ConnectSongs() []table.Row {
	rows, err := music.Songlist(tempPlaylist)
	if err != nil {
		log.Println("Error connecting songs:", err)
		return nil
	}
	return rows
}
