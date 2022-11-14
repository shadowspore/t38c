package t38c

import (
	"context"
	"fmt"
)

// Ping the server.
func (client *Client) Ping(ctx context.Context) error {
	var resp struct {
		Ping string `json:"ping"`
	}

	err := client.jExecute(ctx, &resp, "PING")
	if err != nil {
		return err
	}

	if resp.Ping != "pong" {
		return fmt.Errorf("bad ping response: %v", resp)
	}

	return nil
}

// Health Check
func (client *Client) HealthZ(ctx context.Context) error {
	var response struct {
		OK bool `json:"ok"`
	}

	err := client.jExecute(ctx, &response, "HEALTHZ")
	if err != nil {
		return err
	}

	if !response.OK {
		return fmt.Errorf("health check failed: %v", response)
	}

	return nil
}
