package music

import (
	"Melodex/src/log"
	vlc "github.com/adrg/libvlc-go/v3"
)

func AddSong(songPath string) error {
	if mediaList == nil {
		log.Log.Error("No media list")
	}

	media, err := vlc.NewMediaFromPath(songPath)
	if err != nil {
		return err
	}

	err = mediaList.AddMedia(media)
	if err != nil {
		return err
	}

	log.Log.Infof("Added song to medialist: %s", songPath)
	return nil
}

func PlayMediaList() error {
	err := listPlayer.SetMediaList(mediaList)
	if err != nil {
		return err
	}
	err = listPlayer.Play()
	if err != nil {
		return err
	}
	log.Log.Info("Playing media list")

	return nil
}

func NextSong() error {
	if listPlayer == nil {
		log.Log.Error("No media list")
	}

	err := listPlayer.PlayNext()
	if err != nil {
		return err
	}
	log.Log.Error("Playing next media song")

	return nil
}

func PreviousSong() error {
	if listPlayer == nil {
		log.Log.Error("No media list")
	}

	err := listPlayer.PlayPrevious()
	if err != nil {
		return err
	}
	log.Log.Error("Playing previous media song")

	return nil
}

func PauseSong() error {
	err := listPlayer.TogglePause()
	if err != nil {
		return err
	}
	return nil
}

func StopSong() error {
	err := listPlayer.Stop()
	if err != nil {
		return err
	}
	return nil
}

func MuteSong(option bool) error {
	var err error
	player, err = listPlayer.Player()
	if err != nil {
		return err
	}

	err = player.SetMute(option)
	if err != nil {
		return err
	}
	return nil
}

func SetMediaPosition(factor float32) error {
	var err error
	player, err = listPlayer.Player()
	if err != nil {
		return err
	}

	currentPos, err := getCurrentMediaPosition()

	err = player.SetMediaPosition(currentPos + factor)
	if err != nil {
		return err
	}

	return nil
}

func getCurrentMediaPosition() (float32, error) {
	var err error
	player, err = listPlayer.Player()
	if err != nil {
		return 0, err
	}
	return player.MediaPosition()
}

func SetVolume(volume int) error {
	var err error
	player, err = listPlayer.Player()
	if err != nil {
		return err
	}

	err = player.SetVolume(volume)
	if err != nil {
		return err
	}

	return nil
}

func GetVolume() int {
	var err error
	player, err = listPlayer.Player()
	if err != nil {
		log.Log.Errorf("Error getting volume from player: %v", err)
		return 0
	}

	volume, err := player.Volume()
	if err != nil {
		log.Log.Errorf("Error getting volume from player: %v", err)
		return 0
	}
	return volume
}

//func isMediaListValid() bool {
//	if mediaList == nil {
//		log.Log.Errorf("Media list is not initialized")
//		return false
//	}
//	return true
//}
//
//func playSongAtIndex(index int) {
//	media, err := mediaList.MediaAtIndex(uint(index))
//	if err != nil {
//		log.Log.Errorf("Error getting media at index %d: %v", index, err)
//		return
//	}
//	defer releaseMedia(media, index)
//
//	if err := player.SetMedia(media); err != nil {
//		log.Log.Errorf("Error setting media at index %d: %v", index, err)
//		return
//	}
//
//	if err := player.Play(); err != nil {
//		log.Log.Errorf("Error playing media at index %d: %v", index, err)
//		return
//	}
//
//	log.Log.Infof("Playing song at index %d", index)
//
//	waitForSongCompletion()
//}
//
//func releaseMedia(media *vlc.Media, index int) {
//	if err := media.Release(); err != nil {
//		log.Log.Errorf("Error releasing media at index %d: %v", index, err)
//	}
//}
//
//func waitForSongCompletion() {
//	eventManager, err := player.EventManager()
//	if err != nil {
//		log.Log.Errorf("Error getting media event manager: %v", err)
//		return
//	}
//
//	ch := make(chan struct{})
//
//	go func() {
//		eventManager.Attach(vlc.MediaPlayerEndReached, func(event vlc.Event, data interface{}) {
//			select {
//			case <-ch:
//			default:
//				close(ch)
//			}
//		}, nil)
//	}()
//
//	<-ch
//}
