package Music

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

type Player struct {
	logger        *log.Logger
	context       *oto.Context
	commandChan   chan string
	mu            sync.Mutex
	paused        bool
	currentPlayer *oto.Player
	currentFile   *os.File
	done          chan bool
	stop          chan bool
	stopped       bool
	playlist      []string
	currentIndex  int
}

func NewMusicPlayer(context *oto.Context, logger *log.Logger) *Player {
	return &Player{
		logger:      logger,
		context:     context,
		commandChan: make(chan string),
	}
}

func (mp *Player) PlaySong(directory string, songName string) {
	filePath := filepath.Join(directory, songName)

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	decoder, err := mp3.NewDecoder(f)
	if err != nil {
		panic(err)
	}

	player := mp.context.NewPlayer()
	mp.logger.Printf("player: %#v\n", player)
	mp.currentPlayer = player
	mp.currentFile = f
	mp.done = make(chan bool)
	mp.stop = make(chan bool)

	go func() {
		defer f.Close()
		defer player.Close()
		mp.playAudio(decoder)
	}()

	// mp.waitForCommands()
}

func (mp *Player) playAudio(decoder *mp3.Decoder) {
	buf := make([]byte, 4096)
	for {
		select {
		case <-mp.stop:
			mp.done <- true
			return
		default:
			mp.mu.Lock()
			if mp.paused {
				mp.mu.Unlock()
				continue
			}
			mp.mu.Unlock()
			n, err := decoder.Read(buf)
			if err == io.EOF {
				mp.done <- true
				return
			}
			if err != nil {
				log.Printf("Error reading audio data: %v\n", err)
				mp.done <- true
				return
			}

			// mp.logger.Printf("buf: %d\n", n)

			if n > 0 {
				if _, err := mp.currentPlayer.Write(buf[:n]); err != nil {
					mp.logger.Printf("Error playing audio: %v\n", err)
					mp.done <- true
					return
				} else {
					// mp.logger.Printf("wr: %d\n", nw)
				}
			}
		}
	}
}

func (mp *Player) PauseSong() {
	mp.mu.Lock()
	defer mp.mu.Unlock()
	mp.paused = !mp.paused
}

func (mp *Player) NextSong() {
	mp.stop <- true
	<-mp.done
	mp.currentIndex++
	if mp.currentIndex >= len(mp.playlist) {
		mp.currentIndex = 0
	}
	fmt.Println("Skipping to next track.")
}

func (mp *Player) PreviousSong() {
	mp.stop <- true
	<-mp.done
	mp.currentIndex--
	if mp.currentIndex < 0 {
		mp.currentIndex = len(mp.playlist) - 1
	}
	fmt.Println("Going back to previous track.")
}

func (mp *Player) Stop() {
	mp.stopped = true
	mp.stop <- true
	<-mp.done
	fmt.Println("Playback stopped.")
}
