package music

import (
	"Melodex/src/backend/config"
	"Melodex/src/log"
	vlc "github.com/adrg/libvlc-go/v3"
)

var (
	player *vlc.Player
	media  *vlc.Media
)

func InitVLC() {
	var err error

	if err = vlc.Init("--no-video", "--quiet"); err != nil {
		log.Log.Errorf("Error initializing VLC: %v", err)
	}

	player, err = vlc.NewPlayer()
	if err != nil {
		log.Log.Errorf("Error creating VLC player: %v", err)
	}

	muted, err := player.IsMuted()
	if err != nil {
		log.Log.Errorf("Error checking if player is muted: %v", err)
	} else if muted {
		if err = player.SetMute(false); err != nil {
			log.Log.Errorf("Error unmuting player: %v", err)
		}
	}

	err = player.SetVolume(config.GetVolume())
	log.Log.Infof("Set volume to %d", config.GetVolume())
	if err != nil {
		log.Log.Errorf("Error setting volume to %d: %v", config.GetVolume(), err)
	}
}

func Cleanup() {
	if player != nil {
		err := player.Release()
		if err != nil {
			log.Log.Errorf("Error releasing player: %v", err)
		}
	}

	if media != nil {
		err := vlc.Release()
		if err != nil {
			log.Log.Errorf("Error releasing VLC: %v", err)
		}
	}
}
