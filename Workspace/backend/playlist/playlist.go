package playlist

import (
	"os"
	"strings"
)

var playlistname string

func AddPlaylist(playlistname string) string {
	return playlistname
}

func GetAllSongs() []string {
	var filesWithExt []string
	ext := ".m4a"
	tempPath := "/home/" + os.Getenv("USER") + "/TempSongs"

	files, err := os.ReadDir(tempPath)
	if err != nil {
		return nil
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ext) {
			filesWithExt = append(filesWithExt, file.Name())
		}
	}

	return filesWithExt
}

func ShufflePlaylist(playlist []string) []string {
	shuffeledplaylist := playlist
	return shuffeledplaylist

}

func RemovePlaylist(config string) bool {
	return false
}

func GetPlaylistName(directory string) string {
	return directory
}
