package sharedState

type SharedState struct {
	Searching   bool
	SettingTime bool
	Shuffling   bool
	SongOption  int
	Paused      bool
}

var GlobalState = &SharedState{}

func GetGlobalState() *SharedState {
	return GlobalState
}
