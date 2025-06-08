package music

import (
	"Melodex/src/backend/config"
	"Melodex/src/logger"
	vlc "github.com/adrg/libvlc-go/v3"
)

var (
	player    *vlc.Player
	media     *vlc.Media
	songIndex int
)

// Init initializes the VLC player with configured settings.
// It ensures VLC is not muted and sets the initial volume.
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

// PlaySong plays the provided song file.
// If a song is already playing, it stops the current playback before loading the new song.
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

// PauseSong toggles the playback state between paused and playing.
func PauseSong() {
	if err := player.SetPause(player.IsPlaying()); err != nil {
		logger.Log.Errorf("Error toggling pause: %v", err)
	}
}
func SetMute(option bool) {
	var err error
	if option {
		err = player.SetMute(true)
		if err != nil {
			logger.Log.Errorf("Error toggling mute: %v", err)
		}
	} else {
		err = player.SetMute(false)
		if err != nil {
			logger.Log.Errorf("Error toggling mute: %v", err)
		}
	}
}

func IsMuted() bool {
	state, err := player.IsMuted()
	if err != nil {
		logger.Log.Errorf("Error checking if player is muted: %v", err)
	}
	return state
}

// IsPlaying checks if a song is currently playing.
func IsPlaying() bool {
	return player.IsPlaying()
}

// CurrentSongPosition returns the playback position of the current song as a float.
// If an error occurs, it defaults to 0.0.
func CurrentSongPosition() float32 {
	position, err := player.MediaPosition()
	if err != nil {
		logger.Log.Errorf("Error getting media position: %v", err)
		return 0.0
	}
	return position
}

// SetVolume adjusts the playback volume to the specified level.
func SetVolume(volume int) {
	err := player.SetVolume(volume)
	if err != nil {
		logger.Log.Errorf("Error setting volume to %d: %v", volume, err)
	}
}

// SetMediaPosition moves playback to the specified position in the media.
// Returns false if setting the position fails.
func SetMediaPosition(position float32) bool {
	err := player.SetMediaPosition(position)
	if err != nil {
		logger.Log.Errorf("Error setting media position: %v", err)
		return false
	}
	return true
}

// IncreaseVolume raises the playback volume by the given factor.
// The volume is capped at 100%.
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

// DecreaseVolume lowers the playback volume by the given factor.
// The volume cannot go below 0%.
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

// Stop halts the playback if a song is currently playing.
func Stop() bool {
	if player.IsPlaying() {
		if err := player.Stop(); err != nil {
			logger.Log.Errorf("Error stopping player: %v", err)
			return false
		}
	}
	return true
}

// Cleanup releases all allocated resources related to the VLC player and media.
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
}

func SongLength() int {
	length, err := player.MediaLength()
	if err != nil {
		logger.Log.Errorf("Error getting media length: %v", err)
	}

	return length + 1000
}
