package geofence

import (
	"encoding/json"
	"log"
)

func (client *Client) fenceObjects(req Requestable) (chan Object, error) {
	events, err := client.Fence(req)
	if err != nil {
		return nil, err
	}

	ch := make(chan Object, 10)
	go func() {
		for {
			defer close(ch)
			for event := range events {
				var obj Object
				if err := json.Unmarshal(event, &obj); err != nil {
					log.Printf("bad event: %v", err)
					break
				}

				ch <- obj
			}
		}
	}()

	return ch, nil
}

func (client *Client) fencePoints(req Requestable) (chan Point, error) {
	events, err := client.Fence(req)
	if err != nil {
		return nil, err
	}

	ch := make(chan Point, 10)
	go func() {
		for {
			defer close(ch)
			for event := range events {
				var point Point
				if err := json.Unmarshal(event, &point); err != nil {
					log.Printf("bad event: %v", err)
					break
				}

				ch <- point
			}
		}
	}()
	return ch, nil
}

func (client *Client) fenceBounds(req Requestable) (chan Bounds, error) {
	events, err := client.Fence(req)
	if err != nil {
		return nil, err
	}

	ch := make(chan Bounds, 10)
	go func() {
		for {
			defer close(ch)
			for event := range events {
				var bounds Bounds
				if err := json.Unmarshal(event, &bounds); err != nil {
					log.Printf("bad event: %v", err)
					break
				}

				ch <- bounds
			}
		}
	}()

	return ch, nil
}

func (client *Client) fenceHashes(req Requestable) (chan Hash, error) {
	events, err := client.Fence(req)
	if err != nil {
		return nil, err
	}

	ch := make(chan Hash, 10)
	go func() {
		for {
			defer close(ch)
			for event := range events {
				var hash Hash
				if err := json.Unmarshal(event, &hash); err != nil {
					log.Printf("bad event: %v", err)
					break
				}

				ch <- hash
			}
		}
	}()

	return ch, nil
}

// FenceIntersectsObjects ...
func (client *Client) FenceIntersectsObjects(req *Request) (chan Object, error) {
	req.Cmd = "INTERSECTS"
	return client.fenceObjects(req)
}

// FenceIntersectsPoints ...
func (client *Client) FenceIntersectsPoints(req *Request) (chan Point, error) {
	req.Cmd = "INTERSECTS"
	return client.fencePoints(req)
}

// FenceIntersectsBounds ...
func (client *Client) FenceIntersectsBounds(req *Request) (chan Bounds, error) {
	req.Cmd = "INTERSECTS"
	return client.fenceBounds(req)
}

// FenceIntersectsHashes ...
func (client *Client) FenceIntersectsHashes(req *Request) (chan Hash, error) {
	req.Cmd = "INTERSECTS"
	return client.fenceHashes(req)
}
