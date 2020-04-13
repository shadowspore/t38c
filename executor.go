package t38c

import "context"

// Executor interface
type Executor interface {
	Execute(command string, args ...string) ([]byte, error)
	ExecuteStream(ctx context.Context, command string, args ...string) (chan []byte, error)
}

// ExecutorDialer func
type ExecutorDialer func() (Executor, error)
