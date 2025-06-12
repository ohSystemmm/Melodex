package music

import (
	"Melodex/src/log"
	"fmt"
	vlc "github.com/adrg/libvlc-go/v3"
	"github.com/charmbracelet/bubbles/table"
	"math/rand"
	"time"
)

func PlaySong(songPath string) bool {
	if player.IsPlaying() {
		if err := player.Stop(); err != nil {
			log.Log.Errorf("Error stopping player: %v", err)
			return false
		}
	}

	_, err := player.LoadMediaFromPath(songPath)
	if err != nil {
		log.Log.Errorf("Error loading song: %v", err)
		return false
	}

	songIndex++
	err = player.Play()
	if err = player.Play(); err != nil {
		log.Log.Errorf("Error playing song: %v", err)
		return false
	}
	return true
}
func PauseSong() {
	if err := player.SetPause(player.IsPlaying()); err != nil {
		log.Log.Errorf("Error toggling pause: %v", err)
	}
}
func SetMute(option bool) {
	var err error
	if option {
		err = player.SetMute(true)
		if err != nil {
			log.Log.Errorf("Error toggling mute: %v", err)
		}
	} else {
		err = player.SetMute(false)
		if err != nil {
			log.Log.Errorf("Error toggling mute: %v", err)
		}
	}
}

func IsMuted() bool {
	state, err := player.IsMuted()
	if err != nil {
		log.Log.Errorf("Error checking if player is muted: %v", err)
	}
	return state
}

func IsPlaying() bool {
	return player.IsPlaying()
}

func CurrentSongPosition() float32 {
	position, err := player.MediaPosition()
	if err != nil {
		log.Log.Errorf("Error getting media position: %v", err)
		return 0.0
	}
	return position
}

func SetVolume(volume int) {
	err := player.SetVolume(volume)
	if err != nil {
		log.Log.Errorf("Error setting volume to %d: %v", volume, err)
	}
}

func SetMediaPosition(position float32) bool {
	err := player.SetMediaPosition(position)
	if err != nil {
		log.Log.Errorf("Error setting media position: %v", err)
		return false
	}
	return true
}

func PutVolume(factor int) {
	currentVolume, err := player.Volume()
	if err != nil {
		log.Log.Errorf("Error getting volume: %v", err)
	}

	newVolume := currentVolume + factor
	if newVolume >= 100 {
		newVolume = 100
	} else if newVolume < 0 {
		newVolume = 0
	}

	err = player.SetVolume(newVolume)
	if err != nil {
		log.Log.Errorf("Error setting volume from %d to %d: %v", currentVolume, newVolume, err)
	}
}

func Stop() bool {
	if player.IsPlaying() {
		if err := player.Stop(); err != nil {
			log.Log.Errorf("Error stopping player: %v", err)
			return false
		}
	}
	return true
}

func Cleanup() {
	if player != nil {
		if err := player.Release(); err != nil {
			log.Log.Errorf("Error releasing player: %v", err)
		}
	}

	if media != nil {
		if err := vlc.Release(); err != nil {
			log.Log.Errorf("Error releasing VLC: %v", err)
		}
	}
}

func SongLength() int {
	length, err := player.MediaLength()
	if err != nil {
		log.Log.Errorf("Error getting media length: %v", err)
	}

	return length / 1000
}

var currentIndex int = -1
var songs []table.Row // Stores song path and duration
var path string

func SetPath(pathd string) {
	path = pathd
}

// Set the song list
func SetSongs(songsd []table.Row) {
	songs = songsd
}

// GetSongIndex searches for the songPath and returns its index
func GetSongIndex(songPath string) int {
	for i, row := range songs {
		if row[0] == songPath {
			return i
		}
	}
	return -1 // Song not found
}

// PlayNextSong finds the current song and plays the next one
func PlayNextSong(currentSongPath string) {
	if len(songs) == 0 {
		log.Log.Info("No songs to play")
		return
	}

	index := GetSongIndex(currentSongPath)
	if index == -1 {
		log.Log.Warnf("Song not found: %s", currentSongPath)
		return
	}

	// Move to next song, looping if necessary
	nextIndex := (index + 1) % len(songs)
	Stop()
	PlaySong(songs[nextIndex][0]) // Play the next song
}

func Print() string {
	return songs[1][2]
}

// PlayPreviousSong finds the current song and plays the previous one
func PlayPreviousSong(currentSongPath string) {
	if len(songs) == 0 {
		log.Log.Info("No songs to play")
		return
	}

	index := GetSongIndex(currentSongPath)
	if index == -1 {
		log.Log.Warnf("Song not found: %s", currentSongPath)
		return
	}

	// Move to previous song, looping if necessary
	prevIndex := (index - 1 + len(songs)) % len(songs)
	Stop()
	PlaySong(fmt.Sprintf(songs[prevIndex][0])) // Play the previous song
}

// SelectRandomSong sets `currentIndex` to a random song and plays it
func SelectRandomSong() {
	if len(songs) == 0 {
		log.Log.Info("No songs to play")
		return
	}

	rand.Seed(time.Now().UnixNano())
	currentIndex = rand.Intn(len(songs))
	PlaySong(fmt.Sprintf(songs[currentIndex][0]))
}
