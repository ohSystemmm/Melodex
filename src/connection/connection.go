package connection

import (
	"Melodex/src/backend/config"
	"Melodex/src/backend/music"
	"Melodex/src/backend/playlist"
	"Melodex/src/logger"
	"strings"

	"github.com/charmbracelet/bubbles/table"
)
import (
	"path/filepath"
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
	allSongs, err = music.GenerateSongList(playlistPath, "/home/"+config.GetUser()+"/.cache/melodex")
	if err != nil {
		logger.Log.Errorf("Error connecting to playlist: %v", err)
	}

	return allSongs
}

func Play() {
	playlist.GetAllSongs(allSongs)
	if len(allSongs) == 0 {
		logger.Log.Info("No songs found")
		return
	}
	var song string
	if mode >= 0 {
		for {
			if shuffle {
				song = playlist.GetRandomSong()
			} else {
				if mode == 0 {
					song = playlist.RepeatPlaylist()
				} else {
					song = playlist.RepeatSong()
				}
			}
		}
	} else {
		song = playlist.PlayCurrentSong()
	}
	music.PlaySong(song)
}
