package connection

import (
	"Melodex/src/backend/config"
	"Melodex/src/backend/music"
	"Melodex/src/logger"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/table"
)

var (
	playlistPath string
	currentSong  string
	shuffle      bool
	mode         int
	allSongs     []table.Row
)

func GetPlaylistName() string {
	return filepath.Base(strings.TrimSpace(playlistPath))
}
func SetPlaylistPath(passedPlaylistPath string) {
	playlistPath = passedPlaylistPath
}
func SetCurrentSong(song string) {
	currentSong = song
}
func GetCurrentSong() string {
	return currentSong
}

func SetShuffle(state bool) {
	shuffle = state
}

func SetMode(state int) {
	mode = state
}

func ConnectSongs() []table.Row {
	var err error
	allSongs, err = music.GenerateSongList(playlistPath, "/home/"+config.ConfGetUser()+"/.cache/melodex")
	if err != nil {
		logger.Log.Errorf("Error connecting to playlist: %v", err)
	}

	return allSongs
}

func Play(index int) {
}
