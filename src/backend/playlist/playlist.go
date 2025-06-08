package playlist // TODO Wole Package

import (
	"math/rand"
	"time"

	"github.com/charmbracelet/bubbles/table"
)

var (
	currentIndex int
	songList     []table.Row
)

// GetAllSongs initializes the song list and resets the current index.
func GetAllSongs(songs []table.Row) {
	songList = songs
	currentIndex = 0
}

// IncreaseIndex increments the current song index.
func IncreaseIndex() {
	currentIndex++
}

// DecreaseIndex decrements the current song index.
func DecreaseIndex() {
	currentIndex--
}

// init initializes the random seed to ensure randomized selections.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetRandomSong selects and returns a random song from the list.
// Returns an empty string if the list is empty.
func GetRandomSong() string {
	if len(songList) == 0 {
		return ""
	}
	randomIndex := rand.Intn(len(songList))
	return songList[randomIndex][0]
}

// GetNextSong retrieves the next song in the playlist.
// If the end of the playlist is reached, it loops back to the start.
func GetNextSong() string {
	if len(songList) == 0 {
		return ""
	}
	nextIndex := (currentIndex + 1) % len(songList)
	currentIndex = nextIndex
	return songList[currentIndex][0]
}

// GetPreviousSong retrieves the previous song in the playlist.
// If at the beginning, it loops back to the last song.
func GetPreviousSong() string {
	if len(songList) == 0 {
		return ""
	}
	previousIndex := (currentIndex - 1 + len(songList)) % len(songList)
	currentIndex = previousIndex
	return songList[currentIndex][0]
}

// RepeatPlaylist restarts the playlist from the beginning.
func RepeatPlaylist() string {
	return ""
}

// RepeatSong replays the current song.
func RepeatSong() string {
	return ""
}

// PlayCurrentSong plays the currently selected song.
func PlayCurrentSong() string {
	return ""
}
