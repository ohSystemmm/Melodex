package settings

import (
	"github.com/charmbracelet/bubbles/table"
	"os"
	"path/filepath"
)

type Config struct {
	AppName           string
	Version           string
	Release           string
	SystemUser        string
	ConfigDir         string
	CacheDir          string
	LogDir            string
	ConfigFile        string
	CacheFile         string
	LogFile           string
	PlaylistPath      string
	PlaylistName      string
	IsShuffled        bool
	Mode              int
	CurrentSong       string
	CurrentSongIndex  int
	CurrentSongLength string
	SongList          []table.Row
}

var AppConfig = Config{
	AppName:    "melodex",
	Version:    "0.9.0",
	Release:    "07/01/2025",
	SystemUser: GetUser(),
	ConfigFile: "melodex.toml",
	LogFile:    "melodex.log",
}

func init() {
	AppConfig.ConfigDir = "/home/" + AppConfig.SystemUser + "/.config/melodex/"
	AppConfig.CacheDir = "/home/" + AppConfig.SystemUser + "/.cache/melodex/cache/"
	AppConfig.LogDir = "/home/" + AppConfig.SystemUser + "/.cache/melodex/log/"
	AppConfig.PlaylistPath = "/home/" + AppConfig.SystemUser + "/HolyMoly/008_Music/Best_Songs_Ever-ohSystemmm/"
	AppConfig.PlaylistName = parsePlaylistPath()
}

func GetUser() string {
	envUser := os.Getenv("USER")
	if envUser == "" {
		envUser = "root"
	}
	return envUser
}

func SetConfigFile(configFile string) {
	AppConfig.ConfigFile = configFile
}
func SetPlaylistPath(playlistPath string) {
	AppConfig.PlaylistPath = playlistPath
}
func SetPlaylistName(playlistName string) {
	AppConfig.PlaylistName = playlistName
}
func SetShuffle(shuffle bool) {
	AppConfig.IsShuffled = shuffle
}
func SetMode(mode int) {
	if mode < -1 {
		mode = -1
	} else if mode > 1 {
		mode = 1
	}
	AppConfig.Mode = mode
}
func SetCurrentSong(song string) {
	AppConfig.CurrentSong = song
}
func SetSongList(songList []table.Row) {
	AppConfig.SongList = songList
}

func parsePlaylistPath() string {
	return filepath.Base(AppConfig.PlaylistPath)
}

func SetCurrentSongIndex(index int) {
	AppConfig.CurrentSongIndex = index
}

func IncreaseIndex() {
	if AppConfig.CurrentSongIndex >= len(AppConfig.SongList) {
		AppConfig.CurrentSongIndex = 0
	} else {
		AppConfig.CurrentSongIndex++
	}
}

func DecreaseIndex() {
	if AppConfig.CurrentSongIndex < 0 {
		AppConfig.CurrentSongIndex = len(AppConfig.SongList) - 1
	} else {
		AppConfig.CurrentSongIndex--
	}
}

func IncreaseUpdateCurrentSong() {
	if AppConfig.CurrentSongIndex < len(AppConfig.SongList)-1 {
	} else {
		AppConfig.CurrentSongIndex = 0
	}
	AppConfig.CurrentSong = AppConfig.SongList[AppConfig.CurrentSongIndex][0]
}

func DecreaseUpdateCurrentSong() {
	if AppConfig.CurrentSongIndex > 0 {
	} else {
		AppConfig.CurrentSongIndex = len(AppConfig.SongList) - 1
	}
	AppConfig.CurrentSong = AppConfig.SongList[AppConfig.CurrentSongIndex][0]
}

func SetSongLength(songLength string) {
	AppConfig.CurrentSongLength = songLength
}
