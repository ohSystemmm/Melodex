package Music

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	"Melodex/TUI/SharedState"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

type Player struct {
	sharedState   *SharedState.SharedState
	context       *oto.Context
	currentPlayer *oto.Player
	currentFile   *os.File
	playlist      []string
	currentIndex  int
	mu            sync.Mutex
	pauseCond     *sync.Cond
	// commandChan chan string
}

func NewMusicPlayer(context *oto.Context, newSharedState *SharedState.SharedState) *Player {
	p := &Player{
		sharedState: newSharedState,
		context:     context,
		// commandChan: make(chan string),
	}
	p.pauseCond = sync.NewCond(&p.mu)
	return p
}

func (mp *Player) PlaySong(directory string, songName string) {
	filePath := filepath.Join(directory, songName)
	mp.sharedState.Logger.Printf("Playing %s\n", filePath)

	f, err := os.Open(filePath)
	if err != nil {
		mp.sharedState.Logger.Fatalf("Player: Failed to access Filepath: %v", err)
	}

	decodedMp3, err := mp3.NewDecoder(f)
	if err != nil {
		mp.sharedState.Logger.Fatalf("Player: Failed to create Decoder: %v", err)
	}

	player := mp.context.NewPlayer(decodedMp3)
	mp.currentPlayer = player
	mp.currentFile = f

	go func() {
		defer f.Close()
		defer player.Close()

		mp.currentPlayer.Play()

		for player.IsPlaying() {
			time.Sleep(time.Millisecond)
		}
	}()
}

func (mp *Player) PauseSong() {
	if mp.sharedState.Paused {
		mp.currentPlayer.Play()
		mp.sharedState.Paused = false
	} else {
		mp.currentPlayer.Pause()
		mp.sharedState.Paused = true
	}
}

func (mp *Player) NextSong() {
	mp.currentIndex++
	if mp.currentIndex >= len(mp.playlist) {
		mp.currentIndex = 0
	}
}

func (mp *Player) PreviousSong() {
	mp.currentIndex--
	if mp.currentIndex < 0 {
		mp.currentIndex = len(mp.playlist) - 1
	}
}

func (mp *Player) Stop() {
	if mp.currentPlayer != nil {
		mp.currentPlayer.Close()
	}
	if mp.currentFile != nil {
		mp.currentFile.Close()
	}
}
