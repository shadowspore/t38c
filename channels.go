package t38c

import (
	"context"
)

// Chans returns all channels matching pattern.
func (client *Client) Chans(pattern string) ([]Chan, error) {
	var resp struct {
		Chans []Chan `json:"chans"`
	}

	err := client.jExecute(&resp, "CHANS", pattern)
	if err != nil {
		return nil, err
	}

	return resp.Chans, nil
}

// DelChan remove a specified channel.
func (client *Client) DelChan(name string) error {
	return client.jExecute(nil, "DELCHAN", name)
}

// PDelChan removes all channels that match the specified pattern.
func (client *Client) PDelChan(pattern string) error {
	return client.jExecute(nil, "PDELCHAN", pattern)
}

// PSubscribe subscribes the client to the given patterns.
func (client *Client) PSubscribe(ctx context.Context, handler func(*GeofenceEvent), pattern string) error {
	return client.ExecuteStream(ctx, rawEventHandler(handler), "PSUBSCRIBE", pattern)
}

// SetChan creates a Pub/Sub channel which points to a geofenced search.
// If a channel is already associated to that name, it’ll be overwritten.
// Once the channel is created a client can then listen for events on that channel with SUBSCRIBE or PSUBSCRIBE.
// If expiration less than 0, it will be ignored
func (client *Client) SetChan(name string, query GeofenceQueryBuilder) SetChannelQueryBuilder {
	return newSetChannelQueryBuilder(client, name, query.toCmd())
}

// Subscribe subscribes the client to the specified channels.
func (client *Client) Subscribe(ctx context.Context, handler func(*GeofenceEvent), channels ...string) error {
	return client.ExecuteStream(ctx, rawEventHandler(handler), "SUBSCRIBE", channels...)
}
