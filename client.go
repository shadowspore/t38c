package t38c

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

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
	Server    *Server
}

// Config is a t38c client config.
type Config struct {
	// Tile38 server address.
	//
	// Example: localhost:9851
	Address string

	// Enables debug logging.
	// Executed queries will be printed to stdout.
	Debug bool

	// Allows to perform password authorization.
	Password *string

	// ConnectionPoolSize sets number of connections in the pool.
	// Defaults to 4.
	ConnectionPoolSize int
}

// New creates a new Tile38 client.
func New(cfg Config) (*Client, error) {
	radixPool, err := transport.NewRadix(cfg.Address, cfg.ConnectionPoolSize, cfg.Password)
	if err != nil {
		return nil, err
	}

	return NewWithExecutor(radixPool, cfg.Debug)
}

// NewWithExecutor creates a new Tile38 client with provided executor.
// See Executor interface for more information.
func NewWithExecutor(exec Executor, debug bool) (*Client, error) {
	client := &Client{
		exec:  exec,
		debug: debug,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := client.Ping(ctx); err != nil {
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
	client.Server = &Server{client}

	return client, nil
}

func (client *Client) jExecute(ctx context.Context, resp interface{}, command string, args ...string) error {
	b, err := client.Execute(ctx, command, args...)
	if err != nil {
		return err
	}

	if resp != nil {
		return json.Unmarshal(b, &resp)
	}

	return nil
}

// Execute Tile38 command.
func (client *Client) Execute(ctx context.Context, command string, args ...string) ([]byte, error) {
	resp, err := client.exec.Execute(ctx, command, args...)
	if client.debug {
		log.Printf("\033[34m[%s]\u001B[0m: \u001B[32m%s\u001B[0m", newCmd(command, args...).String(), resp)
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
