package connection

import (
	"github.com/charmbracelet/bubbles/table"
	"path/filepath"
)

var (
	songList     []table.Row
	currentSong  string
	playlistPath string
	shuffleState bool
	modeState    int
	index        int
	playListName string
)

func GetSongList() []table.Row {
	return songList
}
func SetSongList(list []table.Row) {
	songList = list
}

func GetCurrentSong() string {
	return currentSong
}
func SetCurrentSong(song string) {
	currentSong = song
}

func GetPlaylistPath() string {
	return playlistPath
}
func SetPlaylistPath(path string) {
	playlistPath = path
}

func GetShuffleState() bool {
	return shuffleState
}
func SetShuffleState(state bool) {
	shuffleState = state
}

func GetModeState() int {
	return modeState
}
func SetModeState(state int) {
	modeState = state
}

func GetIndex() int {
	return index
}
func IncreaseIndex() {
	index++
}

func GetPlayListName() string {
	return filepath.Base(playlistPath)
}
