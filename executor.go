package t38c

import (
	"context"

	"github.com/zerobounty/tile38-client/transport"
)

var _ Executor = (*transport.RadixPool)(nil)

// Executor represents Tile38 connection.
// Communication should be in JSON format only.
type Executor interface {
	Execute(command string, args ...string) ([]byte, error)
	ExecuteStream(ctx context.Context, handler func([]byte) error, command string, args ...string) error
	Close() error
}

// for internal usage
type tile38Client interface {
	Executor
	jExecute(resp interface{}, command string, args ...string) error
}
