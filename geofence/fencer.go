package geofence

import (
	"context"

	t38c "github.com/lostpeer/tile38-client"
)

// Fencer interface
type Fencer interface {
	Fence(ctx context.Context, cmd t38c.Command) (chan []byte, error)
}

// FencerDialer ...
type FencerDialer func() (Fencer, error)
