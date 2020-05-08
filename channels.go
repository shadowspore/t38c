package t38c

import (
	"context"
	"strconv"
)

// ChanBuilder struct
type ChanBuilder struct {
	Name    string
	Metas   []Meta
	Command Command
	Ex      *int
}

func (ch *ChanBuilder) args() []string {
	var args []string
	args = append(args, ch.Name)

	for _, meta := range ch.Metas {
		args = append(args, "META")
		args = append(args, meta.Name)
		args = append(args, meta.Value)
	}

	if ch.Ex != nil {
		args = append(args, "EX")
		args = append(args, strconv.Itoa(*ch.Ex))
	}

	args = append(args, ch.Command.Name)
	args = append(args, ch.Command.Args...)
	return args
}

// NewChan return new channel builder.
func NewChan(name string, query geofenceQueryBuilder) *ChanBuilder {
	return &ChanBuilder{
		Name:    name,
		Command: query.Cmd(),
	}
}

// Meta ...
// func (ch *ChanBuilder) Meta(name, value string) *ChanBuilder {
// 	ch.Metas = append(ch.Metas, Meta{
// 		Name:  name,
// 		Value: value,
// 	})

// 	return ch
// }

// Expiration set the specified expire time, in seconds.
func (ch *ChanBuilder) Expiration(seconds int) *ChanBuilder {
	ch.Ex = &seconds
	return ch
}

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
func (client *Client) PSubscribe(ctx context.Context, pattern string) (chan GeofenceResponse, error) {
	events, err := client.executor.ExecuteStream(ctx, "PSUBSCRIBE", pattern)
	if err != nil {
		return nil, err
	}

	return unmarshalEvents(events)
}

// SetChan creates a Pub/Sub channel which points to a geofenced search.
// If a channel is already associated to that name, it’ll be overwritten.
// Once the channel is created a client can then listen for events on that channel with SUBSCRIBE or PSUBSCRIBE.
func (client *Client) SetChan(ch *ChanBuilder) error {
	return client.jExecute(nil, "SETCHAN", ch.args()...)
}

// Subscribe subscribes the client to the specified channels.
func (client *Client) Subscribe(ctx context.Context, channels ...string) (chan GeofenceResponse, error) {
	events, err := client.executor.ExecuteStream(ctx, "SUBSCRIBE", channels...)
	if err != nil {
		return nil, err
	}

	return unmarshalEvents(events)
}
