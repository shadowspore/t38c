package t38c

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
)

// ChanBuilder struct
type ChanBuilder struct {
	Name    string
	Metas   []Meta
	Command Command
	Ex      *int
}

// Args ...
func (ch *ChanBuilder) Args() []string {
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

// NewChan ...
func NewChan(name string, req GeofenceRequestable) *ChanBuilder {
	return &ChanBuilder{
		Name:    name,
		Command: req.GeofenceCommand(),
	}
}

// Meta ...
func (ch *ChanBuilder) Meta(name, value string) *ChanBuilder {
	ch.Metas = append(ch.Metas, Meta{
		Name:  name,
		Value: value,
	})

	return ch
}

// Expiration ...
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
func (client *Client) PSubscribe(ctx context.Context, pattern string) (chan Response, error) {
	events, err := client.executor.ExecuteStream(ctx, "PSUBSCRIBE", pattern)
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

// SetChan creates a Pub/Sub channel which points to a geofenced search.
// If a channel is already associated to that name, it’ll be overwritten.
// Once the channel is created a client can then listen for events on that channel with SUBSCRIBE or PSUBSCRIBE.
func (client *Client) SetChan(ch *ChanBuilder) error {
	return client.jExecute(nil, "SETCHAN", ch.Args()...)
}

// Subscribe subscribes the client to the specified channels.
func (client *Client) Subscribe(ctx context.Context, channels ...string) (chan Response, error) {
	events, err := client.executor.ExecuteStream(ctx, "SUBSCRIBE", channels...)
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
