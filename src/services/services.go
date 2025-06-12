package services

import (
	"Melodex/src/backend/caching"
	"Melodex/src/backend/config"
	"Melodex/src/backend/music"
	"Melodex/src/log"
	"github.com/charmbracelet/bubbles/table"
)

func ConnectSongs() []table.Row {
	var err error
	songList, err = caching.GenerateSongList(playlistPath, "/home/"+config.GetUser()+"/.cache/melodex")
	if err != nil {
		log.Log.Errorf("Error connecting to playlist: %v", err)
	}
	music.SetPath(playlistPath)
	music.SetSongs(songList)
	return songList
}

func Play(song string) {
	music.PlaySong(playlistPath + "/" + song)
	currentSong = song
}

func PlayNext() {
	music.PlayNextSong(playlistPath + "/" + currentSong)
}

func PlayPrevious() {
	music.PlayPreviousSong(playlistPath + "/" + currentSong)

}
