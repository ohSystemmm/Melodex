package sharedState

import (
	"log"
)

type SharedState struct {
	Shuffling  bool
	SongOption int
	Searching  bool
	Paused     bool
	Logger     *log.Logger
}

var GlobalState = &SharedState{}

func GetGlobalState() *SharedState {
	return GlobalState
}
