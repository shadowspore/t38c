package t38c

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/mediocregopher/radix/v3"
)

// Tile38Client ...
type Tile38Client struct {
	debug bool
	pool  *radix.Pool
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
	pool, err := radix.NewPool("tcp", address, 10, radix.PoolConnFunc(RadixJSONDialer))
	if err != nil {
		return nil, err
	}

	return NewWithPool(pool, opts...)
}

// NewWithPool ...
func NewWithPool(pool *radix.Pool, opts ...ClientOption) (*Tile38Client, error) {
	client := &Tile38Client{
		pool: pool,
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
func (client *Tile38Client) Execute(command string, args ...string) (resp []byte, err error) {
	err = client.pool.Do(radix.Cmd(&resp, command, args...))
	if client.debug {
		cmd := command
		if len(args) > 0 {
			cmd += " " + strings.Join(args, " ")
		}
		log.Printf("[%s]: %s", cmd, resp)
	}

	return
}

// RadixJSONDialer ...
func RadixJSONDialer(net, addr string) (radix.Conn, error) {
	conn, err := radix.Dial(net, addr)
	if err != nil {
		return nil, err
	}

	var b []byte
	if err := conn.Do(radix.Cmd(&b, "OUTPUT", "json")); err != nil {
		conn.Close()
		return nil, err
	}

	var resp struct {
		Ok bool `json:"ok"`
	}
	if err := json.Unmarshal(b, &resp); err != nil {
		conn.Close()
		return nil, err
	}

	if !resp.Ok {
		conn.Close()
		return nil, fmt.Errorf("bad response: %s", b)
	}
	return conn, nil
}
