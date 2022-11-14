package t38c

import "context"

// Channels struct
type Channels struct {
	client tile38Client
}

// Chans returns all Channels matching pattern.
func (ch *Channels) Chans(ctx context.Context, pattern string) ([]Chan, error) {
	var resp struct {
		Chans []Chan `json:"chans"`
	}

	err := ch.client.jExecute(ctx, &resp, "CHANS", pattern)
	if err != nil {
		return nil, err
	}

	return resp.Chans, nil
}

// DelChan remove a specified channel.
func (ch *Channels) DelChan(ctx context.Context, name string) error {
	return ch.client.jExecute(ctx, nil, "DELCHAN", name)
}

// PDelChan removes all Channels that match the specified pattern.
func (ch *Channels) PDelChan(ctx context.Context, pattern string) error {
	return ch.client.jExecute(ctx, nil, "PDELCHAN", pattern)
}

// PSubscribe subscribes the client to the given patterns.
func (ch *Channels) PSubscribe(ctx context.Context, handler func(*GeofenceEvent) error, pattern string) error {
	return ch.client.ExecuteStream(ctx, rawEventHandler(handler), "PSUBSCRIBE", pattern)
}

// SetChan creates a Pub/Sub channel which points to a geofenced search.
// If a channel is already associated to that name, it’ll be overwritten.
// Once the channel is created a client can then listen for events on that channel with SUBSCRIBE or PSUBSCRIBE.
// If expiration less than 0, it will be ignored
func (ch *Channels) SetChan(name string, query GeofenceQueryBuilder) SetChannelQueryBuilder {
	return newSetChannelQueryBuilder(ch.client, name, query.toCmd())
}

// Subscribe subscribes the client to the specified Channels.
func (ch *Channels) Subscribe(ctx context.Context, handler func(*GeofenceEvent) error, Channels ...string) error {
	return ch.client.ExecuteStream(ctx, rawEventHandler(handler), "SUBSCRIBE", Channels...)
}
