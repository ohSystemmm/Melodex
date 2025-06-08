package sharedState

// SharedState represents the global state of the application,
// tracking various playback and interaction statuses.
type SharedState struct {
	Searching   bool // Indicates whether a search operation is in progress.
	SettingTime bool // Indicates whether the user is setting the playback time.
	Shuffling   bool // Indicates whether shuffle mode is active.
	SongOption  int  // Stores the selected song option.
	Paused      bool // Tracks whether playback is currently paused.
}

var GlobalState = &SharedState{}

// GetGlobalState returns a pointer to the global shared state,
// allowing access and modification across the application.
func GetGlobalState() *SharedState {
	return GlobalState
}
