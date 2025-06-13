package music

import (
	"Melodex/src/settings"
)

func PlayNextSong() {
	settings.IncreaseIndex()
	PlaySong(settings.AppConfig.PlaylistPath + settings.AppConfig.SongList[settings.AppConfig.CurrentSongIndex][0])
}

func PlayPreviousSong() {
	settings.DecreaseIndex()
	PlaySong(settings.AppConfig.PlaylistPath + settings.AppConfig.SongList[settings.AppConfig.CurrentSongIndex][0])
}

func SelectRandomSong() {

}
