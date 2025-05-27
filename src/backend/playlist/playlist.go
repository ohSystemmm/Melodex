package playlist

import (
	"fmt"
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

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomSong() string {
	if len(songList) == 0 {
		return ""
	}
	return fmt.Sprintf("%v", songList[rand.Intn(len(songList))][0])
}

func GetNextSong() string {
	if len(songList) == 0 {
		return ""
	}
	song := fmt.Sprintf("%v", songList[currentIndex%len(songList)][0])
	currentIndex++
	return song
}
