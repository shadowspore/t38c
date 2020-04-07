package t38c

// Executor interface
type Executor interface {
	Execute(command string, args ...string) ([]byte, error)
}

// ExecutorDialer func
type ExecutorDialer func() (Executor, error)

// FenceExecutor interface
type FenceExecutor interface {
	Fence(command string, args ...string) (FenceChan, error)
}

// FenceChan ...
type FenceChan chan FenceEvent

// FenceEvent struct
type FenceEvent struct {
	Data []byte
	Err  error
}

// FenceExecutorDialer func
type FenceExecutorDialer func() (FenceExecutor, error)
