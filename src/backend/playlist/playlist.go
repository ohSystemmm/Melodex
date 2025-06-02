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
	if len(songList) == 0 {
		return ""
	}
	randomIndex := rand.Intn(len(songList))
	return songList[randomIndex][0]
}

func GetNextSong() string {
	if len(songList) == 0 {
		return ""
	}
	nextIndex := (currentIndex + 1) % len(songList)
	currentIndex = nextIndex
	return songList[currentIndex][0]
}

func GetPreviousSong() string {
	if len(songList) == 0 {
		return ""
	}
	previousIndex := (currentIndex - 1 + len(songList)) % len(songList)
	currentIndex = previousIndex
	return songList[currentIndex][0]
}

func RepeatPlaylist() string {
	return ""
}

func RepeatSong() string {
	return ""
}

func PlayCurrentSong() string {
	return ""
}
