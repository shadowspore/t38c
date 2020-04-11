package geofence

import (
	"context"
	"encoding/json"
	"log"
)

// Fence ...
func (client *Client) Fence(ctx context.Context, req Requestable) (chan Response, error) {
	events, err := client.execRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	ch := make(chan Response, 10)
	go func() {
		defer close(ch)
		for event := range events {
			var resp Response
			if err := json.Unmarshal(event, &resp); err != nil {
				log.Printf("bad event: %v", err)
				break
			}

			ch <- resp
		}
	}()

	return ch, nil
}
