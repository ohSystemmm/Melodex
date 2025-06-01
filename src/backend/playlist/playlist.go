package playlist

import (
	"math/rand"
	"time"

	"github.com/charmbracelet/bubbles/table"
)

var (
	currentIndex int
	songList     []table.Row
)

func GetAllSongs(songs []table.Row) {
	songList = songs
	currentIndex = 0
}

func IncreaseIndex() {
	currentIndex++
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomSong() string {
	// Return a Random song from songList
	return ""
}

func GetNextSong() string {
	// Return Selected Song  + Index 1
	return ""
}

func GetPreviousSong() string {
	// Return Selected Song -1
	return ""
}
