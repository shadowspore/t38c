package t38c

import "context"

type Server struct {
	client tile38Client
}

// WARNING: This erases all data in Tile38 DB!
func (sv *Server) FlushDB(ctx context.Context) error {
	var resp struct{}

	err := sv.client.jExecute(ctx, &resp, "FLUSHDB")
	// Explicit is better than implicit.
	if err != nil {
		return err
	}

	return nil
}
