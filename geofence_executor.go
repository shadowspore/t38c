package t38c

// GeofenceExecutor interface
type GeofenceExecutor interface {
	Fence(command string, args ...string) (GeofenceChan, error)
}

// GeofenceChan ...
type GeofenceChan chan GeofenceEvent

// GeofenceEvent struct
type GeofenceEvent []byte

// GeofenceExecutorDialer func
// type GeofenceExecutorDialer func() (GeofenceExecutor, error)
