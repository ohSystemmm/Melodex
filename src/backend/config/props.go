package config

var GlobalConfig *Config

func GetUser() string {
	return GlobalConfig.General.User
}
func GetSetUp() string {
	return GlobalConfig.General.SetUp
}
func GetDefaultVolume() string {
	return GlobalConfig.General.DefaultVolume
}
func GetDefaultPlaylist() string {
	return GlobalConfig.General.DefaultPlaylist
}
func GetPlaylists() []string {
	return GlobalConfig.Playlist.Playlists
}
func GetBorderColor() string {
	return GlobalConfig.Design.BorderColor
}
func GetForegroundColor() string {
	return GlobalConfig.Design.ForegroundColor
}
func GetBackgroundColor() string {
	return GlobalConfig.Design.BackgroundColor
}
