package Backend

import (
	"log"

	vlc "github.com/adrg/libvlc-go/v3"
)

func Run() {
	if err := vlc.Init(); err != nil {
		log.Fatal(err)
	}
	defer vlc.Release()

	player, err := vlc.NewPlayer()
	if err != nil {
		log.Fatal(err)
	}
	defer player.Release()

}
