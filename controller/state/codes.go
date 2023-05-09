package state

type State uint8

const (
	STATE_INITIALIZING = State(iota)
	STATE_INITIALIZED
	STATE_RUNNING
	STATE_ERROR
	STATE_TERMINATING
)

func (sv State) String() string {
	switch sv {
	case STATE_INITIALIZING:
		return "INITIALIZING"
	case STATE_INITIALIZED:
		return "INITALIZED"
	case STATE_RUNNING:
		return "RUNNING"
	case STATE_ERROR:
		return "ERROR"
	case STATE_TERMINATING:
		return "TERMINATING"
	default:
		return "Unknown State"
	}
}
