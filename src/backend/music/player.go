package music

import (
	"Melodex/src/log"
)

func PlaySong(songPath string) {
	if player.IsPlaying() {
		if err := player.Stop(); err != nil {
			log.Log.Errorf("Error stopping player: %v", err)
			return
		}
	}

	_, err := player.LoadMediaFromPath(songPath)
	if err != nil {
		log.Log.Errorf("Error loading song: %v", err)
		return
	}

	err = player.Play()
	if err = player.Play(); err != nil {
		log.Log.Errorf("Error playing song: %v", err)
		return
	}
}

func PauseSong() {
	err := player.SetPause(player.IsPlaying())
	if err != nil {
		log.Log.Errorf("Error toggling pause: %v", err)
	}
}

func IsPlaying() bool {
	return player.IsPlaying()
}

func Stop() {
	if player.IsPlaying() {
		err := player.Stop()
		if err != nil {
			log.Log.Errorf("Error stopping player: %v", err)
		}
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

func CurrentMediaPosition() float32 {
	position, err := player.MediaPosition()
	if err != nil {
		log.Log.Errorf("Error getting media position: %v", err)
		return 0.0
	}
	return position
}

func SetMediaPosition(position float32) {
	err := player.SetMediaPosition(position)
	if err != nil {
		log.Log.Errorf("Error setting media position: %v", err)
	}
}

func ChangeVolumeBy(factor int) {
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

func SongLength() int {
	length, err := player.MediaLength()
	if err != nil {
		log.Log.Errorf("Error getting media length: %v", err)
	}

	return length / 1000
}
