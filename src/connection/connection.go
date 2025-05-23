package connection

import (
	"Melodex/src/backend/music"
	"log"

	"github.com/charmbracelet/bubbles/table"
)

// Services but name is cooler

// var tempPlaylist = "/home/" + os.Getenv("USER") + "/Music/"
// var tempPlaylist = "/home/" + os.Getenv("USER") + "/HolyMoly/008_Music/Best_Songs_Ever-ohSystemmm/"
var (
	playlistPath string
	currentSong  string
)

func SetPlaylistPath(passedPlaylistPath string) {
	playlistPath = passedPlaylistPath
}
func GetPlaylistPath() string {
	return playlistPath + "/"
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
	return rows
}
