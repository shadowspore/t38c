package t38c

import (
	"fmt"
)

// Ping the server.
func (client *Tile38Client) Ping() error {
	var resp struct {
		Ping string `json:"ping"`
	}

	if err := client.execute("PING", &resp); err != nil {
		return fmt.Errorf("ping: %v", err)
	}

	if resp.Ping != "pong" {
		return fmt.Errorf("bad ping response: %v", resp)
	}

	return nil
}
