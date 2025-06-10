package music

import (
	"Melodex/src/log"
	vlc "github.com/adrg/libvlc-go/v3"
)

var (
	player     *vlc.Player
	mediaList  *vlc.MediaList
	listPlayer *vlc.ListPlayer
	songIndex  int
	mode       int
	shuffle    bool
)

/* NOTE mode
* -1 = No Repeat
*  0 = Repeat Playlist
*  1 = Repeat Song
 */

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

	listPlayer, err = vlc.NewListPlayer()
	if err != nil {
		log.Log.Errorf("Error creating list player: %v", err)
	}

	songIndex = 0
	log.Log.Infof("VLC initialized successfully!")
}

func CleanUp() error {
	var err error

	err = player.Release()
	if err != nil {
		return err
	}

	err = mediaList.Release()
	if err != nil {
		return err
	}

	err = listPlayer.Release()
	if err != nil {
		return err
	}

	err = vlc.Release()
	if err != nil {
		return err
	}

	log.Log.Info("VLC cleaned up successfully!")
	return nil
}
