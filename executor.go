package t38c

// Executor interface
type Executor interface {
	Execute(command string, args ...string) ([]byte, error)
}

// ExecutorDialer func
type ExecutorDialer func() (Executor, error)
