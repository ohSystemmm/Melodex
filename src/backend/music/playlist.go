package music

import (
	"Melodex/src/log"
	vlc "github.com/adrg/libvlc-go/v3"
	"math/rand"
	"time"
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

func PlayPlaylistFixedOrder() {
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

	for i := 0; i < counter; i++ {
		media, err := mediaList.MediaAtIndex(uint(i))
		if err != nil {
			log.Log.Errorf("Error getting media at index %d: %v", i, err)
			continue
		}

		err = player.SetMedia(media)
		if err != nil {
			log.Log.Errorf("Error setting media at index %d: %v", i, err)
			continue
		}

		err = player.Play()
		if err != nil {
			log.Log.Errorf("Error playing media at index %d: %v", i, err)
			continue
		}

		log.Log.Infof("Playing song at index %d", i)

		for player.IsPlaying() {
			time.Sleep(500 * time.Millisecond)
		}
		err = media.Release()
		if err != nil {
			log.Log.Errorf("Error releasing media at index %d: %v", i, err)
		}
	}
}

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

		for player.IsPlaying() {
			time.Sleep(500 * time.Millisecond)
		}
		err = media.Release()
		if err != nil {
			log.Log.Errorf("Error releasing media at index %d: %v", randomIndex, err)
		}
	}
}
