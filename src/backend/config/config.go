package config

type Config struct {
	General  General  `toml:"general"`
	Playlist Playlist `toml:"playlist"`
	Design   Design   `toml:"design"`
}

type General struct {
	User            string `toml:"user"`
	SetUp           string `toml:"setup"`
	DefaultVolume   string `toml:"default_volume"`
	DefaultPlaylist string `toml:"default_playlist"`
}

type Playlist struct {
	Playlists []string `toml:"playlists"`
}

type Design struct {
	BorderColor      string `toml:"border"`
	ForegroundColor  string `toml:"foreground"`
	BackgroundColor  string `toml:"background"`
	VolumeBarColor   string `toml:"volume_bar"`
	MusicSliderColor string `toml:"music_slider"`
}
