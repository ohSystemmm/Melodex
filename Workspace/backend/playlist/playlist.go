package playlist

var playlistname string

func AddPlaylist(playlistname string) string {
	return playlistname
}

func GetPlaylist() string {
	return playlistname
}

func GetAllSongs(direcory string) []string {
	return []string{}
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
