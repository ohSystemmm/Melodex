package sharedState

import (
	"log"
)

type SharedState struct {
	Searching   bool
	SettingTime bool
	Shuffling   bool
	SongOption  int
	Paused      bool
	Logger      *log.Logger
}

var GlobalState = &SharedState{}

func GetGlobalState() *SharedState {
	return GlobalState
}
