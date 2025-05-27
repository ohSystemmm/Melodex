package sharedState

import (
	"log"
)

type SharedState struct {
	Shuffling bool
	Looping   bool
	// Loop Playlist
	LoopPL    bool
	Searching bool
	Paused    bool
	Logger    *log.Logger
}

var GlobalState = &SharedState{}

func GetGlobalState() *SharedState {
	return GlobalState
}
