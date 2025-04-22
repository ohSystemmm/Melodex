package connection

import (
	"Melodex/src/backend/music"
	"Melodex/src/backend/playlist"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/table"
)

// Services but name is cooler

func ConnectSongs() []table.Row {
	var rows []table.Row

	songs := playlist.GetAllSongs()

	for _, song := range songs {
		songPath := "/home/" + os.Getenv("USER") + "/TempSongs/" + song

		length, err := music.SongLength("", songPath)
		if err != nil {
			log.Println("Error getting length for song:", song, err)
			continue
		}

		minutes := length / 60
		seconds := length % 60

		formattedLength := fmt.Sprintf("%02d:%02d", minutes, seconds)

		row := table.Row{
			song,
			formattedLength,
		}

		rows = append(rows, row)
	}

	return rows
}
