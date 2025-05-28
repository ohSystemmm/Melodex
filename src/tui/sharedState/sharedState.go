package sharedState

import (
	"log"
)

type SharedState struct {
	Searching   bool
	SettingTime bool
	Paused      bool
	Logger      *log.Logger
}

var GlobalState = &SharedState{}

func GetGlobalState() *SharedState {
	return GlobalState
}
