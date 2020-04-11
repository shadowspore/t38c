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

	}
	ch, err := client.fencer.Fence(cmd.Name, cmd.Args...)
	if client.debug {
		if err != nil {
			log.Printf("geofence request: [%s]: %v", cmd, err)
			return nil, err
		}

		log.Printf("geofence request: [%s]: ok", cmd)
		proxyCh := make(chan []byte, 10)
		go func() {
			defer close(proxyCh)

			for event := range ch {
				log.Printf("[event]: %s", event)
				proxyCh <- event
			}
		}()
		return proxyCh, nil
	}

	return ch, err
}
