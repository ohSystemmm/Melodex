package music

import (
	"Melodex/src/log"
	vlc "github.com/adrg/libvlc-go/v3"
)

var (
	player    *vlc.Player
	mediaList *vlc.MediaList
	songIndex int
)

func InitVLC() {
	var err error

	if err = vlc.Init("--no-video", "--quiet"); err != nil {
		log.Log.Errorf("Error initializing VLC: %v", err)
		return
	}

	player, err = vlc.NewPlayer()
	if err != nil {
		log.Log.Errorf("Error creating VLC player: %v", err)
		return
	}

	mediaList, err = vlc.NewMediaList()
	if err != nil {
		log.Log.Errorf("Error creating media list: %v", err)
		return
	}

	songIndex = 0
	log.Log.Infof("VLC initialized successfully!")
}
