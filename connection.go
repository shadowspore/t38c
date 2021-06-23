package t38c

import (
	"fmt"
)

// Ping the server.
func (client *Client) Ping() error {
	var resp struct {
		Ping string `json:"ping"`
	}

	err := client.jExecute(&resp, "PING")
	if err != nil {
		return err
	}

	if resp.Ping != "pong" {
		return fmt.Errorf("bad ping response: %v", resp)
	}

	return nil
}

// Health Check
func (client *Client) HealthZ() error {
	var response struct {
		OK bool `json:"ok"`
	}

	err := client.jExecute(&response, "HEALTHZ")
	if err != nil {
		return err
	}

	if !response.OK {
		return fmt.Errorf("health check failed: %v", response)
	}

	return nil
}
