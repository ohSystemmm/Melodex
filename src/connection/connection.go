package connection

import (
	"Melodex/src/backend/music"
	"Melodex/src/backend/playlist"
	"log"

	"github.com/charmbracelet/bubbles/table"
)

// Services but name is cooler

// var tempPlaylist = "/home/" + os.Getenv("USER") + "/Music/"
// var tempPlaylist = "/home/" + os.Getenv("USER") + "/HolyMoly/008_Music/Best_Songs_Ever-ohSystemmm/"
var (
	playlistPath string
	currentSong  string
	shuffle      bool
	mode         int

	allSongs []table.Row
)

func GetShuffle() bool {
	return shuffle
}
func SetShuffle(state bool) {
	shuffle = state
}
func GetMode() int {
	return mode
}
func SetMode(newMode int) {
	mode = newMode
}
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
	allSongs = rows
	return rows
}

func Play() {
	playlist.GetAllSongs(allSongs)
	if mode < 0 {
		music.PlaySong(playlistPath + "/" + currentSong)
	} else if mode == 0 {
		if shuffle {
			music.PlaySong(playlistPath + "/" + playlist.GetRandomSong())
		} else {
			music.PlaySong(playlistPath + "/" + playlist.GetNextSong())
		}
	} else {
		for {
			music.PlaySong(currentSong)
		}
	}
}
