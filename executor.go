package t38c

import (
	"context"

	"github.com/sythang/t38c/transport"
)

var _ Executor = (*transport.Radix)(nil)

// Executor represents Tile38 connection.
// Communication should be in JSON format only.
type Executor interface {
	Execute(ctx context.Context, command string, args ...string) ([]byte, error)
	ExecuteStream(ctx context.Context, handler func([]byte) error, command string, args ...string) error
	Close() error
}

// for internal usage
type tile38Client interface {
	Executor
	jExecute(ctx context.Context, resp interface{}, command string, args ...string) error
}
