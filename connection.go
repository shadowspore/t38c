package t38c

import (
	"encoding/json"
	"fmt"
)

// Ping the server.
func (client *Tile38Client) Ping() error {
	var resp struct {
		Ping string `json:"ping"`
	}

	b, err := client.Execute("PING")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return err
	}

	if resp.Ping != "pong" {
		return fmt.Errorf("bad ping response: %v", resp)
	}

	return nil
}
