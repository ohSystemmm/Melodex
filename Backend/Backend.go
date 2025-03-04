package Backend

import (
	"fmt"
	"log"

	"Melodex/Backend/Music"
	"github.com/hajimehoshi/oto"
)

func Run() {
	sourceDir := "./Backend/Music/source_songs"
	destDir := "./Backend/Music/mp3_songs"

	mp3Files, err := Music.ConvertToMP3(sourceDir, destDir)
	if err != nil {
		log.Fatalf("Error getting music files: %v", err)
	}

	if len(mp3Files) == 0 {
		fmt.Println("No files were converted.")
		return
	}

	context, err := oto.NewContext(44100, 2, 2, 65536)
	if err != nil {
		log.Fatalf("Failed to create audio context: %v", err)
	}
	defer context.Close()

	musicPlayer := Music.NewMusicPlayer(context)
	musicPlayer.PlaySongs(mp3Files)
}
