package Music

import (
	"log"
	"time"

	vlc "github.com/adrg/libvlc-go/v3"
)

var (
	player *vlc.Player
	media  *vlc.Media
)

// Creates a new VLC player instance
func Init() {
	if err := vlc.Init("--no-video", "--quiet"); err != nil {
		log.Fatal(err)
	}

	var err error
	player, err = vlc.NewPlayer()
	if err != nil {
		log.Fatal(err)
	}
}

// loads and plays songs
func PlaySong(song string) {
	var err error
	media, err = player.LoadMediaFromPath(song)
	if err != nil {
		log.Fatal(err)
	}

	err = player.Play()
	if err != nil {
		log.Fatal(err)
	}
}

// toggles pause and play
func PauseSong() {
	player.SetPause(!player.IsPlaying())
}

// Metadata stuff, not implemented yet, to lazy rn lmfao
func Metadata() {
}

// returns songlength and position
func SongLength() (length int, position float32, err error) {
	length, err = player.MediaLength()
	if err != nil {
		log.Println("Error getting length:", err)
		return
	}

	position, err = player.MediaPosition()
	if err != nil {
		log.Println("Error getting position:", err)
		return
	}

	return length, position, nil
}

// sets the volume (0–100)
func SetVolume(volume int) {
	if err := player.SetVolume(volume); err != nil {
		log.Println("Failed to set volume:", err)
	}
}

// toggles mute
func SetMute() {
	isMuted, err := player.IsMuted()
	if err != nil {
		log.Println("Failed to get mute status:", err)
		return
	}

	err = player.SetMute(!isMuted)
	if err != nil {
		log.Println("Failed to toggle mute:", err)
	}
}

// sets playback position (0.0–1.0) - float32 btw
func SetMediaPosition(pos float32) {
	if err := player.SetMediaPosition(pos); err != nil {
		log.Println("Failed to set position:", err)
	}
}

// Sleeps and stops playback after time is up
func Sleep(duration time.Duration) {
	log.Printf("Sleeping for %s...\n", duration)
	time.Sleep(duration)

	log.Println("Sleep complete. Stopping playback.")
	Stop()
}

// stops playback
func Stop() {
	if player != nil {
		player.Stop()
	}
	if media != nil {
		media.Release()
	}
}

// cleanup to release resources
func Cleanup() {
	if player != nil {
		player.Release()
	}
	vlc.Release()
}
