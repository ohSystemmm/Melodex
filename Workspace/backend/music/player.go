package music

import (
	"log"
	"time"

	vlc "github.com/adrg/libvlc-go/v3"
)

type VlcPlayer struct {
	player *vlc.Player
	media  *vlc.Media
	logger *log.Logger
}

var (
	v VlcPlayer
)

func SetLogger(newLogger *log.Logger) {
	v.logger = newLogger
}

// Creates a new VLC player instance
func Init() {
	if err := vlc.Init("--no-video", "--quiet"); err != nil {
		v.logger.Println(err)
	}

	var err error
	v.player, err = vlc.NewPlayer()
	if err != nil {
		v.logger.Println(err)
	}
}

// loads and plays songs
func PlaySong(song string) {
	var err error
	v.media, err = v.player.LoadMediaFromPath(song)
	if err != nil {
		v.logger.Println(err)
	}

	err = v.player.Play()
	if err != nil {
		v.logger.Println(err)
	}
}

// Returns wether the Player is Playing
func IsPlaying() bool {
	return v.player.IsPlaying()
}

// toggles pause and play
func PauseSong() {
	err := v.player.SetPause(IsPlaying())
	if err != nil {
		v.logger.Println(err)
	}
}

// Metadata stuff, not implemented yet, to lazy rn lmfao
func Metadata() {
}

// returns song length
// NOTE the purpose of directory is questionable
func SongLength(directory string, song string) (length int, err error) {
	if v.media == nil {
		v.media, err = v.player.LoadMediaFromPath(song)
		if err != nil {
			v.logger.Println("Error loading media:", err)
			return 0, err
		}
		v.logger.Println("Media loaded successfully.")
	}

	duration, err := v.media.Duration()
	if err != nil {
		v.logger.Println("Error getting media duration:", err)
		return 0, err
	}

	return int(duration / 1000), nil
}

func SongPosition() (position float32, err error) {
	position, err = v.player.MediaPosition()
	if err != nil {
		v.logger.Println("Error getting position:", err)
		return
	}

	return position, nil
}

// sets the volume (0–100)
func SetVolume(volume int) {
	if err := v.player.SetVolume(volume); err != nil {
		v.logger.Println("Failed to set volume:", err)
	}
}

// toggles mute
func SetMute() {
	isMuted, err := v.player.IsMuted()
	if err != nil {
		v.logger.Println("Failed to get mute status:", err)
		return
	}

	err = v.player.SetMute(!isMuted)
	if err != nil {
		v.logger.Println("Failed to toggle mute:", err)
	}
}

// sets playback position (0.0–1.0) - float32 btw
func SetMediaPosition(pos float32) {
	if err := v.player.SetMediaPosition(pos); err != nil {
		v.logger.Println("Failed to set position:", err)
	}
}

// Sleeps and stops playback after time is up
func Sleep(duration time.Duration) {
	v.logger.Printf("Sleeping for %s...\n", duration)
	time.Sleep(duration)

	log.Println("Sleep complete. Stopping playback.")
	Stop()
}

// Stops playback
func Stop() {
	if v.player != nil {
		v.player.Stop()
	}
	if v.media != nil {
		v.media.Release()
	}
}

// Cleanup to release resources
func (v VlcPlayer) Cleanup() {
	if v.player != nil {
		v.player.Release()
	}
	vlc.Release()
}
