package music

import (
	"Melodex/src/log"
	"errors"
	vlc "github.com/adrg/libvlc-go/v3"
)

func AddSong(songPath string) error {
	if mediaList == nil {
		log.Log.Error("No media list")
		return errors.New("no media list")
	}

	media, err := vlc.NewMediaFromPath(songPath)
	if err != nil {
		return err
	}

	err = mediaList.AddMedia(media)
	if err != nil {
		return err
	}

	return nil
}

func PlayMediaList() error {
	if mediaList == nil {
		log.Log.Error("No media list")
		return errors.New("no media list")
	}

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
		log.Log.Error("No list player initialized")
		return errors.New("no list player")
	}

	err := listPlayer.PlayNext()
	if err != nil {
		return err
	}
	log.Log.Info("Playing next media song")

	return nil
}

func PreviousSong() error {
	if listPlayer == nil {
		log.Log.Error("No list player initialized")
		return errors.New("no list player")
	}

	err := listPlayer.PlayPrevious()
	if err != nil {
		return err
	}
	log.Log.Info("Playing previous media song")

	return nil
}

func PauseSong() error {
	if listPlayer == nil {
		log.Log.Error("No list player initialized")
		return errors.New("no list player")
	}

	err := listPlayer.TogglePause()
	if err != nil {
		return err
	}
	return nil
}
func StopSong() error {
	if listPlayer == nil {
		log.Log.Error("No list player initialized")
		return errors.New("no list player")
	}

	err := listPlayer.Stop()
	if err != nil {
		return err
	}
	return nil
}

// --
func GetActivePlayer() (*vlc.Player, error) {
	if listPlayer == nil {
		log.Log.Error("No list player initialized")
		return nil, errors.New("no list player")
	}

	player, err := listPlayer.Player()
	if err != nil {
		log.Log.Errorf("Error retrieving active player: %v", err)
		return nil, err
	}
	return player, nil
}

func SetMediaPosition(factor float32) error {
	player, err := GetActivePlayer()
	if err != nil {
		return err
	}

	currentPos, err := player.MediaPosition()
	if err != nil {
		return err
	}

	err = player.SetMediaPosition(currentPos + factor)
	if err != nil {
		return err
	}

	return nil
}

func SetVolume(volume int) error {
	player, err := GetActivePlayer()
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
	player, err := GetActivePlayer()
	if err != nil {
		log.Log.Errorf("Error retrieving active player: %v", err)
		return 0
	}

	volume, err := player.Volume()
	if err != nil {
		log.Log.Errorf("Error getting volume: %v", err)
		return 0
	}
	return volume
}

func IsMuted() bool {
	player, err := GetActivePlayer()
	if err != nil {
		log.Log.Errorf("Error retrieving active player: %v", err)
		return false
	}

	state, err := player.IsMuted()
	if err != nil {
		log.Log.Errorf("Error retrieving mute state: %v", err)
		return false
	}

	return state
}

func SetMute(option bool) error {
	player, err := GetActivePlayer()
	if err != nil {
		log.Log.Errorf("Error retrieving active player: %v", err)
		return err
	}

	err = player.SetMute(option)
	if err != nil {
		log.Log.Errorf("Error setting mute: %v", err)
		return err
	}

	return nil
} // --
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

func PutVolume(factor int) {
	player, err := GetActivePlayer()
	if err != nil {
		log.Log.Errorf("Error retrieving active player: %v", err)
	}

	vol := GetCurrentVolume() + factor
	if vol < 0 {
		vol = 0
	} else if vol > 100 {
		vol = 100
	}

	if err := player.SetVolume(vol); err != nil {
		log.Log.Errorf("Error setting volume: %v", err)
	}
	log.Log.Infof("Setting volume to %d", vol)
}

func GetCurrentVolume() int {
	player, err := GetActivePlayer()
	if err != nil {
		log.Log.Errorf("Error retrieving active player: %v", err)
		return 0
	}
	volume, err := player.Volume()
	if err != nil {
		log.Log.Errorf("Error getting volume from player: %v", err)
	}
	return volume
}

func getCurrentMediaPosition() (float32, error) {
	var err error
	player, err = listPlayer.Player()
	if err != nil {
		return 0, err
	}
	return player.MediaPosition()
}

func IsPlaying() bool {
	var err error
	player, err = listPlayer.Player()
	if err != nil {
		log.Log.Errorf("Error getting player from player: %v", err)
		return false
	}
	if player.IsPlaying() {
		return true
	}
	return false
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
