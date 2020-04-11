package geofence

// Fencer interface
type Fencer interface {
	Fence(command string, args ...string) (chan []byte, error)
}

// FencerDialer ...
type FencerDialer func() (Fencer, error)
