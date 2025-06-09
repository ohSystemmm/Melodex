package music

import (
	"Melodex/src/log"
	vlc "github.com/adrg/libvlc-go/v3"
)

func AddSong(songPath string) {
	if mediaList == nil {
		log.Log.Errorf("Medialist not initialized")
	}

	media, err := vlc.NewMediaFromPath(songPath)
	if err != nil {
		log.Log.Errorf("Error loading song: %v", err)
	}

	err = mediaList.AddMedia(media)
	if err != nil {
		log.Log.Errorf("Error adding song to medialist: %v", err)
	}

	log.Log.Infof("Added song to medialist: %s", songPath)
}

func isMediaListValid() bool {
	if mediaList == nil {
		log.Log.Errorf("Media list is not initialized")
		return false
	}
	return true
}

func playSongAtIndex(index int) {
	media, err := mediaList.MediaAtIndex(uint(index))
	if err != nil {
		log.Log.Errorf("Error getting media at index %d: %v", index, err)
		return
	}
	defer releaseMedia(media, index)

	if err := player.SetMedia(media); err != nil {
		log.Log.Errorf("Error setting media at index %d: %v", index, err)
		return
	}

	if err := player.Play(); err != nil {
		log.Log.Errorf("Error playing media at index %d: %v", index, err)
		return
	}

	log.Log.Infof("Playing song at index %d", index)

	waitForSongCompletion()
}

func releaseMedia(media *vlc.Media, index int) {
	if err := media.Release(); err != nil {
		log.Log.Errorf("Error releasing media at index %d: %v", index, err)
	}
}

func waitForSongCompletion() {
	eventManager, err := player.EventManager()
	if err != nil {
		log.Log.Errorf("Error getting media event manager: %v", err)
		return
	}

	ch := make(chan struct{})

	go func() {
		eventManager.Attach(vlc.MediaPlayerEndReached, func(event vlc.Event, data interface{}) {
			select {
			case <-ch:
			default:
				close(ch)
			}
		}, nil)
	}()

	<-ch
}
