package t38c

import (
	"fmt"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

// Tile38Client ...
type Tile38Client struct {
	addr  string
	debug bool
	conn  redis.Conn
}

// ClientOption ...
type ClientOption func(*Tile38Client)

// Debug ...
func Debug() ClientOption {
	return func(c *Tile38Client) {
		c.debug = true
	}
}

// New ...
func New(address string, opts ...ClientOption) (*Tile38Client, error) {
	conn, err := redis.Dial("tcp", address,
		redis.DialConnectTimeout(time.Second*10),
	)
	if err != nil {
		return nil, err
	}
	client := &Tile38Client{
		addr: address,
		conn: conn,
	}

	for _, opt := range opts {
		opt(client)
	}

	if _, err := client.Execute("OUTPUT", "json"); err != nil {
		return nil, err
	}

	if err := client.Ping(); err != nil {
		return nil, err
	}

	return client, nil
}

// Execute command
func (client *Tile38Client) Execute(command string, args ...interface{}) ([]byte, error) {
	resp, err := client.conn.Do(command, args...)
	if client.debug {
		log.Printf("[%s]: %s", command, resp)
	}

	if err != nil {
		return nil, err
	}

	body, ok := resp.([]byte)
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", resp)
	}

	return body, nil
}
