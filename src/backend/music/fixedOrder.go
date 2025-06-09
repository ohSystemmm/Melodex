package music

import (
	"Melodex/src/log"
)

func PlayPlaylistFixedOrder() {
	if !isMediaListValid() {
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
		playSongAtIndex(i)
	}
}
