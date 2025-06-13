package config

type Config struct {
	General  General  `toml:"general"`
	Playlist Playlist `toml:"playlist"`
	Design   Design   `toml:"design"`
}

type General struct {
	User            string `toml:"user"`
	DefaultPlaylist string `toml:"default_playlist"`
	DefaultVolume   string `toml:"default_volume"`
}

type Playlist struct {
	Playlists []string `toml:"playlists"`
}

type Design struct {
	Border     string `toml:"border"`
	Foreground string `toml:"foreground"`
	Background string `toml:"background"`
}

var (
	loadedConfig   *Config
	configFilePath string
)

func InitConfig() {
	configFilePath = "/home/ohsystemmm/.config/melodex/melodex.toml"
	loadedConfig = LoadConfig()
}
