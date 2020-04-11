package geofence

import (
	"encoding/json"
	"log"
)

func (client *Client) Fence(req Requestable) (chan Response, error) {
	events, err := client.execRequest(req)
	if err != nil {
		return nil, err
	}

	ch := make(chan Response, 10)
	go func() {
		for {
			defer close(ch)
			for event := range events {
				var resp Response
				if err := json.Unmarshal(event, &resp); err != nil {
					log.Printf("bad event: %v", err)
					break
				}

				ch <- resp
			}
		}
	}()

	return ch, nil
}

