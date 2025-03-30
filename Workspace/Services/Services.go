package Services

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/ebitengine/oto/v3"
	"log"

	"Melodex/Backend/Music"
	"Melodex/TUI/SharedState"
)

// TODO: these shouldn't be global, put them in a struct
var (
	selectedSong string
	otoContext   *oto.Context
	directory    string
	sharedState  *SharedState.SharedState
)

// NewService needs to be called once, and only once
func NewService(newSharedState *SharedState.SharedState) {
	sharedState = newSharedState
	var err error
	var readyChan chan struct{}
	op := &oto.NewContextOptions{}
	op.SampleRate = 44100
	op.ChannelCount = 2
	op.Format = oto.FormatSignedInt16LE
	otoContext, readyChan, err = oto.NewContext(op)
	if err != nil {
		sharedState.Logger.Fatalf("Creating oto context failed")
	}
	<-readyChan
}

func GetSongList(newdirectory string) []table.Row {
	directory = newdirectory
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

func PlaySelectedSong() *Music.Player {
	if otoContext == nil {
		log.Fatalf("NewService must be called before using PlaySelectedSong")
	}

	musicPlayer := Music.NewMusicPlayer(otoContext, sharedState)
	go func() {
		musicPlayer.PlaySong(directory, selectedSong)
	}()
	return musicPlayer
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
