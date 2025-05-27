package playlist

import (
	"fmt" // Import "fmt" for string formatting
	"math/rand"
	"time"

	"github.com/charmbracelet/bubbles/table"
)

var (
	currentIndex int
	songList     []table.Row // each row: {Songname, duration}
)

// GetAllSongs initializes the songList with the provided songs.
func GetAllSongs(songs []table.Row) {
	songList = songs
	currentIndex = 0 // Reset current index when new songs are loaded
}

func init() {
	rand.Seed(time.Now().UnixNano()) // Seeding only once for better randomness
}

// GetRandomSong returns a random song from the songList as a string.
// It returns an empty string if the songList is empty.
func GetRandomSong() string {
	if len(songList) == 0 {
		return ""
	}
	randomIndex := rand.Intn(len(songList))
	// Assuming the first element of the row is the song name
	return fmt.Sprintf("%v", songList[randomIndex][0])
}

// GetNextSong returns the next song in sequence from the songList as a string.
// It cycles back to the beginning if the end of the list is reached.
// It returns an empty string if the songList is empty.
func GetNextSong() string {
	if len(songList) == 0 {
		return ""
	}
	if currentIndex >= len(songList) {
		currentIndex = 0 // Cycle back to the beginning
	}
	song := fmt.Sprintf("%v", songList[currentIndex][0]) // Assuming the first element is the song name
	currentIndex++
	return song
}
