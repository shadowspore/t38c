package geofence

import "log"

// Client struct
type Client struct {
	debug bool
	exec  Executor
}

// New ...
func New(dialer ExecutorDialer, debug bool) (*Client, error) {
	exec, err := dialer()
	if err != nil {
		return nil, err
	}

	client := &Client{
		exec:  exec,
		debug: debug,
	}

	return client, nil
}

// Fence ...
func (client *Client) Fence(req Requestable) (chan []byte, error) {
	cmd := req.GeofenceCommand()
	if client.debug {
		log.Printf("geofence request: [%s %s]", cmd.Name, cmd.Args)
	}
	ch, err := client.exec.Fence(cmd.Name, cmd.Args...)
	return ch, err
}
