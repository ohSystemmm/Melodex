package music

import (
	"Melodex/src/backend/config"
	"Melodex/src/logger"
	vlc "github.com/adrg/libvlc-go/v3"
	"time"
)

var (
	player    *vlc.Player
	media     *vlc.Media
	songIndex int
)

// FUNCTION: Initializes VLC
func Init() {
	var err error

	if err = vlc.Init("--no-video", "--quiet"); err != nil {
		logger.Log.Errorf("Error initializing VLC: %v", err)
	}

	player, err = vlc.NewPlayer()
	if err != nil {
		logger.Log.Errorf("Error creating VLC player: %v", err)
	}

	muted, err := player.IsMuted()
	if err != nil {
		logger.Log.Errorf("Error checking if player is muted: %v", err)
	} else if muted {
		if err = player.SetMute(false); err != nil {
			logger.Log.Errorf("Error unmuting player: %v", err)
		}
	}

	err = player.SetVolume(config.ConfGetVolume())
	logger.Log.Infof("Set volume to %d", config.ConfGetVolume())
	if err != nil {
		logger.Log.Errorf("Error setting volume to %d: %v", config.ConfGetVolume(), err)
	}
	songIndex = 0
}

// FUNCTION: Plays the provided song
func PlaySong(songPath string) bool {
	if player.IsPlaying() {
		if err := player.Stop(); err != nil {
			logger.Log.Errorf("Error stopping player: %v", err)
			return false
		}
	}

	_, err := player.LoadMediaFromPath(songPath)
	if err != nil {
		logger.Log.Errorf("Error loading song: %v", err)
		return false
	}

	songIndex++
	err = player.Play()
	if err = player.Play(); err != nil {
		logger.Log.Errorf("Error playing song: %v", err)
		return false
	}
	return true
}

// FUNCTION: Pauses song
func PauseSong() {
	if err := player.SetPause(player.IsPlaying()); err != nil {
		logger.Log.Errorf("Error toggling pause: %v", err)
	}
}

// FUNCTION: Returns the current state of the song
func IsPlaying() bool {
	return player.IsPlaying()
}

// FUNCTION: Gets the current Song Position
func CurrentSongPosition() float32 {
	position, err := player.MediaPosition()
	if err != nil {
		logger.Log.Errorf("Error getting media position: %v", err)
		return 0.0
	}
	return position
}

// FUNCTION: Sets the volume
func SetVolume(volume int) {
	err := player.SetVolume(volume)
	if err != nil {
		logger.Log.Errorf("Error setting volume to %d: %v", volume, err)
	}
}

// FUNCTION: Sets Media Position
func SetMediaPosition(position float32) bool {
	err := player.SetMediaPosition(position)
	if err != nil {
		logger.Log.Errorf("Error setting media position: %v", err)
		return false
	}
	return true
}

// FUNCTION: Increases volume by params value
func IncreaseVolume(factor int) {
	currentVolume, err := player.Volume()
	if err != nil {
		logger.Log.Errorf("Error getting volume: %v", err)
	}

	newVolume := currentVolume + factor
	if newVolume >= 100 {
		newVolume = 100
	}

	err = player.SetVolume(newVolume)
	if err != nil {
		logger.Log.Errorf("Error increasing volume from %d to %d: %v", currentVolume, newVolume, err)
	}
}

// FUNCTION: Decreases volume by params value
func DecreaseVolume(factor int) {
	currentVolume, err := player.Volume()
	if err != nil {
		logger.Log.Errorf("Error getting volume: %v", err)
	}

	newVolume := currentVolume - factor
	if newVolume <= 0 {
		newVolume = 0
	}

	err = player.SetVolume(newVolume)
	if err != nil {
		logger.Log.Errorf("Error decreasing volume from %d to %d: %v", currentVolume, newVolume, err)
	}
}

// FUNCTION: Stops the player
func Stop() bool {
	if player.IsPlaying() {
		if err := player.Stop(); err != nil {
			logger.Log.Errorf("Error stopping player: %v", err)
			return false
		}
	}
	return true
}

// FUNCTION: Releases all resources
func Cleanup() {
	if player != nil {
		if err := player.Release(); err != nil {
			logger.Log.Errorf("Error releasing player: %v", err)
		}
	}

	if media != nil {
		if err := vlc.Release(); err != nil {
			logger.Log.Errorf("Error releasing VLC: %v", err)
		}
	}
	return true
}

func WaitTillSongEnd() {
	for {
		currentPosition, err := player.MediaPosition()
		if err != nil {
			logger.Log.Errorf("Error getting media position: %v", err)
			return
		}

		mediaLength, err := player.MediaLength()
		if err != nil {
			logger.Log.Errorf("Error getting media length: %v", err)
			return
		}

		if currentPosition >= float32(mediaLength) {
			break
		}

		time.Sleep(500 * time.Millisecond)
	}

	logger.Log.Info("Song Finished")
}

func GetIndex() int {
	return songIndex
}

func GetCurrentSongLength() float64 {
	if player == nil {
		return 0.0
	}
	length, err := player.MediaLength()
	if err != nil {
		logger.Log.Errorf("Error getting media length: %v", err)
	}
	return float64(length)
}
