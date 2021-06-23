package t38c

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/tidwall/gjson"
	"github.com/xjem/t38c/transport"
)

// Client allows you to interact with the Tile38 server.
type Client struct {
	debug bool
	exec  Executor

	Search    *Search
	Keys      *Keys
	Webhooks  *Hooks
	Channels  *Channels
	Scripting *Scripting
	Geofence  *Geofence
}

type clientParams struct {
	debug    bool
	password *string
	poolSize int
}

// ClientOption ...
type ClientOption func(*clientParams)

// Debug option.
var Debug = ClientOption(func(c *clientParams) {
	c.debug = true
})

// WithPassword option.
func WithPassword(password string) ClientOption {
	return func(c *clientParams) {
		c.password = &password
	}
}

// SetPoolSize option.
func SetPoolSize(size int) ClientOption {
	return func(c *clientParams) {
		c.poolSize = size
	}
}

// New creates a new Tile38 client.
// By default uses redis pool with 5 connections.
// In debug mode will also print commands which will be sent to the server.
func New(addr string, opts ...ClientOption) (*Client, error) {
	params := &clientParams{poolSize: 5}
	for _, opt := range opts {
		opt(params)
	}

	radixPool, err := transport.NewRadixPool(addr, params.poolSize, params.password)
	if err != nil {
		return nil, err
	}

	return NewWithExecutor(radixPool, params.debug)
}

// NewWithExecutor creates a new Tile38 client with provided executor.
// See Executor interface for more information.
func NewWithExecutor(exec Executor, debug bool) (*Client, error) {
	client := &Client{
		exec:  exec,
		debug: debug,
	}

	if err := client.Ping(); err != nil {
		return nil, err
	}

	// Health check can be used this way to test the readiness of tile38
	// if healthError := client.HealthZ(); healthError != nil {
	// 	return nil, healthError
	// }

	client.Webhooks = &Hooks{client}
	client.Geofence = &Geofence{client}
	client.Keys = &Keys{client}
	client.Search = &Search{client}
	client.Scripting = &Scripting{client}
	client.Channels = &Channels{client}

	return client, nil
}

func (client *Client) jExecute(resp interface{}, command string, args ...string) error {
	b, err := client.Execute(command, args...)
	if err != nil {
		return err
	}

	if resp != nil {
		return json.Unmarshal(b, &resp)
	}

	return nil
}

// Execute Tile38 command.
func (client *Client) Execute(command string, args ...string) ([]byte, error) {
	resp, err := client.exec.Execute(command, args...)
	if client.debug {
		log.Printf("[%s]: %s", newCmd(command, args...).String(), resp)
	}

	if err != nil {
		return nil, err
	}

	if !gjson.GetBytes(resp, "ok").Bool() {
		return nil, fmt.Errorf(gjson.GetBytes(resp, "err").String())
	}

	return resp, nil
}

// ExecuteStream used for Tile38 commands with streaming response.
func (client *Client) ExecuteStream(ctx context.Context, handler func([]byte) error, command string, args ...string) error {
	if client.debug {
		log.Printf("[%s]", newCmd(command, args...).String())
	}

	return client.exec.ExecuteStream(ctx, handler, command, args...)
}

// Close closes all connections in the pool and rejects future execution calls.
// Blocks until all streams are closed.
//
// NOTE: custom Executor implementation may change behavior.
func (client *Client) Close() error {
	return client.exec.Close()
}
