package Music

import (
	"log"
	"os"
	"path/filepath"
	// "sync"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

type Player struct {
	logger      *log.Logger
	context     *oto.Context
	commandChan chan string
	// pauseCond     *sync.Cond
	// mu            sync.Mutex
	currentPlayer *oto.Player
	currentFile   *os.File
	playlist      []string
	paused        bool
	currentIndex  int
	// remainingSong int
}

func NewMusicPlayer(context *oto.Context, logger *log.Logger) *Player {
	p := &Player{
		logger:      logger,
		context:     context,
		commandChan: make(chan string),
	}
	// p.pauseCond = sync.NewCond(&p.mu)
	return p
}

func (mp *Player) PlaySong(directory string, songName string) {
	filePath := filepath.Join(directory, songName)
	mp.logger.Printf("Playing %s\n", filePath)

	f, err := os.Open(filePath)
	if err != nil {
		mp.logger.Fatalf("Player: Failed to access Filepath")
	}

	decoded, err := mp3.NewDecoder(f)
	if err != nil {
		mp.logger.Fatalf("Player: Failed to create Decoder")
	}

	player := mp.context.NewPlayer(decoded)
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

// func (mp *Player) playAudio(decoder *mp3.Decoder) {
// 	buf := make([]byte, 1024)
// 	for {
// 		select {
// 		// case <-mp.stop:

// 			mp.done <- true
// 			return
// 		default:
// 			mp.mu.Lock()
// 			// Handle pause state
// 			if mp.paused {
// 				mp.pauseCond.Wait()
// 			}
// 			mp.mu.Unlock()

// 			// Audio processing
// 			n, err := decoder.Read(buf)
// 			if err == io.EOF {
// 				mp.done <- true
// 				return
// 			}
// 			if err != nil {
// 				mp.logger.Printf("Read error: %v", err)
// 				mp.done <- true
// 				return
// 			}

// 			if _, err := mp.currentPlayer.Write(buf[:n]); err != nil {
// 				mp.logger.Printf("Write error: %v", err)
// 				mp.done <- true
// 				return
// 			} else {
// 				// The n is form the prefious if, replace the underscore for testing
// 				// mp.logger.Printf("Wrote %d bytes\n", n)
// 				// time.Sleep(time.Duration(n) * time.Nanosecond)
// 			}
// 		}
// 	}
// }

// BUG PauseSong does not Resume the Song, only end it
func (mp *Player) PauseSong() {
	// mp.mu.Lock()
	// defer mp.mu.Unlock()

	// mp.paused = !mp.paused
	// if !mp.paused {
	// 	mp.pauseCond.Broadcast()
	// }
	if mp.paused {
		mp.currentPlayer.Play()
		mp.paused = false
		// mp.pauseCond.Broadcast()
		// mp.logger.Printf("To this point it works")
	} else {
		// mp.remainingSong = mp.currentPlayer.BufferedSize()
		mp.currentPlayer.Pause()
		mp.paused = true
		// mp.pauseCond.Broadcast()
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
	err := mp.currentPlayer.Close()
	if err != nil {
		mp.logger.Printf("Closing the Player failed")
	}
}
