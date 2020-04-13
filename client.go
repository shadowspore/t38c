package t38c

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tidwall/gjson"
)

// Client struct
type Client struct {
	debug    bool
	executor Executor
}

// New ...
func New(addr string, debug bool) (*Client, error) {
	dialer := NewRadixPool(addr, 5)
	return NewWithDialer(dialer, debug)
}

// NewWithDialer ...
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

// ExecuteCmd ...
func (client *Client) ExecuteCmd(cmd Command) ([]byte, error) {
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

// Execute command
func (client *Client) Execute(command string, args ...string) ([]byte, error) {
	return client.ExecuteCmd(NewCommand(command, args...))
}

// JExecute ...
func (client *Client) JExecute(resp interface{}, command string, args ...string) error {
	b, err := client.Execute(command, args...)
	if err != nil {
		return err
	}

	if resp != nil {
		return json.Unmarshal(b, &resp)
	}

	return nil
}
