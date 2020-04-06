package t38c

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/mediocregopher/radix/v3"
	"github.com/tidwall/gjson"
)

// Tile38Client struct
type Tile38Client struct {
	addr  string
	debug bool
	pool  *radix.Pool
}

// ClientOptions struct
type ClientOptions struct {
	Debug    bool
	Addr     string
	PoolSize int
	Pool     *radix.Pool
}

// New ...
func New(ops ClientOptions) (client *Tile38Client, err error) {
	var pool *radix.Pool
	if ops.Pool == nil {
		pool, err = radix.NewPool("tcp", ops.Addr, ops.PoolSize, radix.PoolConnFunc(RadixJSONDialer))
		if err != nil {
			return nil, err
		}
	}

	client = &Tile38Client{
		debug: ops.Debug,
		pool:  pool,
	}

	if err := client.Ping(); err != nil {
		return nil, err
	}

	return client, nil
}

// Execute command
func (client *Tile38Client) Execute(command string, args ...string) ([]byte, error) {
	var resp []byte
	err := client.pool.Do(radix.Cmd(&resp, command, args...))
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
