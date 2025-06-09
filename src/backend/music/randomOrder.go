package music

import (
	"Melodex/src/log"
	vlc "github.com/adrg/libvlc-go/v3"
	"math/rand"
	"time"
)

func PlayPlaylistRandomOrder() {
	if mediaList == nil {
		log.Log.Errorf("Media list is not initialized")
		return
	}

	counter, err := mediaList.Count()
	if err != nil {
		log.Log.Errorf("Error counting media list: %v", err)
		return
	}

	if counter == 0 {
		log.Log.Errorf("Media list is empty")
		return
	}

	rand.Seed(time.Now().UnixNano())
	indices := rand.Perm(counter)

	for _, randomIndex := range indices {
		media, err := mediaList.MediaAtIndex(uint(randomIndex))
		if err != nil {
			log.Log.Errorf("Error getting media at index %d: %v", randomIndex, err)
			continue
		}

		err = player.SetMedia(media)
		if err != nil {
			log.Log.Errorf("Error setting media at index %d: %v", randomIndex, err)
			continue
		}
		err = player.Play()
		if err != nil {
			log.Log.Errorf("Error playing media at index %d: %v", randomIndex, err)
			continue
		}

		log.Log.Infof("Playing song at index %d", randomIndex)

		eventManager, err := player.EventManager()
		if err != nil {
			log.Log.Errorf("Error getting media event manager: %v", err)
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

		err = media.Release()
		if err != nil {
			log.Log.Errorf("Error releasing media at index %d: %v", randomIndex, err)
		}
	}
}
