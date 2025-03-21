package Services

import (
	"Melodex/Backend/Music"
	"github.com/charmbracelet/bubbles/table"
	"github.com/hajimehoshi/oto"
	"log"
)

var selectedSong string

func GetSongList(directory string) []table.Row {
	files, err := Music.GetFilesWithExtensions(directory)
	if err != nil {
		return []table.Row{
			{"Error Reading Directory " + directory, "--:--"},
		}
	}

	var rows []table.Row
	for name, path := range files {
		rows = append(rows, table.Row{name, Music.GetDuration(path)})
	}

	return rows
}

func PlaySelectedSong(directory string, song string) {

	context, err := oto.NewContext(44100, 2, 2, 65536)
	if err != nil {
		log.Fatalf("Failed to create audio context: %v", err)
	}
	defer context.Close()

	musicPlayer := Music.NewMusicPlayer(context)
	musicPlayer.PlaySong(directory, song)
}

func SetSelectedSong(song string) {
	selectedSong = song
}

func GetSelectedSong() string {
	return selectedSong
}

func GetSongLength(song string) int {
	return 0
}

//func PauseSong() func {
//	return
//}
//
//func ShuffleSong() func {
//	return
//}
//
//func RepeatSong() func {
//	return
//}
//
//func RepeatPlaylist() func {
//	return
//}
