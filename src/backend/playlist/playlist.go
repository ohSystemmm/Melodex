package playlist

import (
	"math/rand"
	"time"

	"github.com/charmbracelet/bubbles/table"
)

var (
	songNames []table.Row
)

// Extract song names from playlist
func parsePlaylist(playlist []table.Row) {
	songNames = make([]table.Row, len(playlist))

	for i, row := range playlist {
		if len(row) > 0 {
			songNames[i] = table.Row{row[0]}
		}
	}
}

func playlistNormal() []table.Row {
	return songNames
}

// Shuffle the songNames slice randomly
func playlistShuffled() []table.Row {
	if len(songNames) == 0 {
		return nil
	}

	shuffled := make([]table.Row, len(songNames))
	copy(shuffled, songNames)

	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}
