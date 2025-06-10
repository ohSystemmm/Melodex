package services

var (
	playlistName string
	currentSong  string
)

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
