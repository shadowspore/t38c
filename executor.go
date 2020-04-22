package t38c

import "context"

// Executor represents Tile38 connection.
// Communication should be in JSON format only.
type Executor interface {
	Execute(command string, args ...string) ([]byte, error)
	ExecuteStream(ctx context.Context, command string, args ...string) (chan []byte, error)
}

// ExecutorDialer is a function which creates Executor instance.
type ExecutorDialer func(password *string) (Executor, error)
