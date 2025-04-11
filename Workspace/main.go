package main

import (
	"fmt"
	"log"
	"os"

	vlc "github.com/adrg/libvlc-go/v3"
)

func main() {
	if err := vlc.Init(); err != nil {
		log.Fatal(err)
	}
	defer vlc.Release()

	player, err := vlc.NewPlayer()
	if err != nil {
		log.Fatal(err)
	}
	defer player.Release()

	player.SetMute(false)
	media, err := player.LoadMediaFromPath("Backend/Music/source_songs/Song3.wav")
	if err != nil {
		log.Fatal("Failed to load media:", err)
	}
	defer media.Release()

	if err := player.Play(); err != nil {
		log.Fatal("Failed to play:", err)
	}

	select {}
}

func greeter() {
	file, err := os.ReadFile("Assets/Greeter.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n" + string(file) + "\n")
}
