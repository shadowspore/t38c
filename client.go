package t38c

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/tidwall/gjson"
)

// Client allows you to interact with the Tile38 server.
type Client struct {
	debug    bool
	executor Executor
}

// New creates a new Tile38 client.
// By default uses redis pool with 5 connections.
// In debug mode will also print commands which will be sent to the server.
func New(addr string, debug bool) (*Client, error) {
	dialer := NewRadixPool(addr, 5)
	return NewWithDialer(dialer, debug)
}

// NewWithDialer creates a new Tile38 client with provided dialer.
// See Executor interface for more information.
func NewWithDialer(dialer ExecutorDialer, debug bool) (*Client, error) {
	executor, err := dialer()
	if err != nil {
		return nil, err
	}

	client := &Client{
		debug:    debug,
		executor: executor,
	}

	if err := client.Ping(); err != nil {
		return nil, err
	}

	return client, nil
}

func (client *Client) executeCmd(cmd Command) ([]byte, error) {
	resp, err := client.executor.Execute(cmd.Name, cmd.Args...)
	if client.debug {
		log.Printf("[%s]: %s", cmd, resp)
	}

	if err != nil {
		return nil, err
	}

	if !gjson.GetBytes(resp, "ok").Bool() {
		return nil, fmt.Errorf("command: %s: %s", cmd, gjson.GetBytes(resp, "err").String())
	}

	return resp, nil
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
	return client.executeCmd(NewCommand(command, args...))
}

// ExecuteStream used for Tile38 commands with streaming response.
func (client *Client) ExecuteStream(ctx context.Context, command string, args ...string) (chan []byte, error) {
	return client.executor.ExecuteStream(ctx, command, args...)
}
