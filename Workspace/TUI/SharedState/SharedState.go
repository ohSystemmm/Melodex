package SharedState

type SharedState struct {
	Searching bool
}

var GlobalState = &SharedState{}

func GetGlobalState() *SharedState {
	return GlobalState
}
