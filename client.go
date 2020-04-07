package t38c

import (
	"fmt"
	"log"
	"strings"

	"github.com/tidwall/gjson"
)

// Tile38Client struct
type Tile38Client struct {
	debug    bool
	executor Executor
}

// ClientOption func
type ClientOption func(*Tile38Client)

// Debug option
var Debug = func(c *Tile38Client) {
	c.debug = true
}

// New ...
func New(dialer ExecutorDialer, opts ...ClientOption) (*Tile38Client, error) {
	executor, err := dialer()
	if err != nil {
		return nil, err
	}

	client := &Tile38Client{
		executor: executor,
	}

	for _, opt := range opts {
		opt(client)
	}

	if err := client.Ping(); err != nil {
		return nil, err
	}

	return client, nil
}

// Execute command
func (client *Tile38Client) Execute(command string, args ...string) ([]byte, error) {
	resp, err := client.executor.Execute(command, args...)
	if client.debug {
		cmd := command
		if len(args) > 0 {
			cmd += " " + strings.Join(args, " ")
		}
		log.Printf("[%s]: %s", cmd, resp)
	}

	if err != nil {
		return nil, err
	}

	if !gjson.GetBytes(resp, "ok").Bool() {
		cmd := command
		if len(args) > 0 {
			cmd += " " + strings.Join(args, " ")
		}

		return nil, fmt.Errorf("command '%s': %s", cmd, gjson.GetBytes(resp, "err").String())
	}

	return resp, nil
}
