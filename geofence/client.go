package geofence

import "log"

// Client struct
type Client struct {
	debug  bool
	fencer Fencer
}

// New ...
func New(dialer FencerDialer, debug bool) (*Client, error) {
	fencer, err := dialer()
	if err != nil {
		return nil, err
	}

	client := &Client{
		fencer: fencer,
		debug:  debug,
	}

	return client, nil
}

// Fence ...
func (client *Client) Fence(req Requestable) (chan []byte, error) {
	cmd := req.GeofenceCommand()
	if client.debug {
		log.Printf("geofence request: [%s %s]", cmd.Name, cmd.Args)
	}
	ch, err := client.fencer.Fence(cmd.Name, cmd.Args...)
	return ch, err
}
