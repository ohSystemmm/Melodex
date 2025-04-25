package connection

import (
	// "Melodex/backend/playlist"
	// "fmt"

	"github.com/charmbracelet/bubbles/table"
)

// Services but name is cooler

// func ConnectSongs() []table.Row {
// 	var rows []table.Row

// 	songs := playlist.GetAllSongs()

// 	for _, song := range songs {
// 		songPath := "/home/" + os.Getenv("USER") + "/TempSongs/" + song

// 		length, err := music.SongLength("", songPath)
// 		if err != nil {
// 			log.Println("Error getting length for song:", song, err)
// 			continue
// 		}

// 		minutes := length / 60
// 		seconds := length % 60

// 		// Format the length as "mm:ss"
// 		formattedLength := fmt.Sprintf("%02d:%02d", minutes, seconds)

// 		row := table.Row{
// 			song,
// 			formattedLength,
// 		}

// 		rows = append(rows, row)
// 	}

// 	return rows
// }

func ConnectSongs() []table.Row {
	var rows []table.Row
	// durations, err := playlist.DurationOfAllSongs("/path/to/songs")
	// if err != nil {
	// 	// handle error
	// }
	// for song, duration := range durations {
	// 	minutes := duration / 60
	// 	seconds := duration % 60

	// 	// Format the duration as "mm:ss"
	// 	formattedDuration := fmt.Sprintf("%02d:%02d", minutes, seconds)
	// 	row := table.Row{
	// 		song,
	// 		formattedDuration,
	// 	}

	// 	rows = append(rows, row)
	// }

	return rows
}
