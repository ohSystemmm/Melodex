package services

import (
	"Melodex/src/backend/cache"
	"Melodex/src/backend/music"
	"Melodex/src/log"
	"Melodex/src/settings"
	"github.com/charmbracelet/bubbles/table"
	"path/filepath"
)

var (
	playlistName = "PlaylistName" // TODO
	currentSong  string
	songList     []table.Row
)

func AddAllSongsToMediaList() {
	var err error
	for _, song := range songList {
		err = music.AddSong(song[0])
		if err != nil {
			log.Log.Errorf("Error while adding song to music: %s", err)
			continue
		}
	}
	log.Log.Infof("Added all songs to music")
}

func search(currentSong string) int {
	for i, row := range songList {
		if row[0] == currentSong {
			return i
		}
	}
	return -1
}

func SetPlaylistName(name string) {
	playlistName = "Playlist: " + name
}
func GetPlaylistName() string {
	return playlistName
}

func SetCurrentSong(songName string) {
	currentSong = songName
}
func GetCurrentSong() string {
	return currentSong
}

func SongList() []table.Row {
	var err error
	songList, err = cache.GenerateSongList("/home/ohsystemmm/HolyMoly/008_Music/Best_Songs_Ever-ohSystemmm", settings.AppConfig.CacheDir)
	if err != nil {
		log.Log.Errorf("Error connecting to playlist: %v", err)
	}

	var parsedSongList []table.Row
	for _, row := range songList {
		parsedSongList = append(parsedSongList, table.Row{filepath.Base(row[0]), row[1]})
	}

	AddAllSongsToMediaList()
	return parsedSongList
}
