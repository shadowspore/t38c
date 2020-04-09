package geofence

// Executor interface
type Executor interface {
	Fence(command string, args ...string) (chan []byte, error)
}
