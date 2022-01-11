package t38c

type Server struct {
	client tile38Client
}

// WARNING: This erases all data in Tile38 DB!
func (sv *Server) FlushDB() error {
	var resp struct{}

	err := sv.client.jExecute(&resp, "FLUSHDB")

	// Explicit is better than implicit.
	if err != nil {
		return err
	}

	return nil
}
