package t38c

import (
	"encoding/json"
	"fmt"

	"github.com/mediocregopher/radix/v3"
)

var _ Executor = (*RadixPoolExecutor)(nil)

// RadixPoolExecutor struct
type RadixPoolExecutor struct {
	pool *radix.Pool
}

// NewRadixPool ...
func NewRadixPool(addr string, size int) ExecutorDialer {
	return func() (Executor, error) {
		pool, err := radix.NewPool("tcp", addr, size, radix.PoolConnFunc(RadixJSONDialer))
		if err != nil {
			return nil, err
		}

		return &RadixPoolExecutor{
			pool: pool,
		}, nil
	}
}

// Execute command
func (rad *RadixPoolExecutor) Execute(command string, args ...string) ([]byte, error) {
	var resp []byte
	err := rad.pool.Do(radix.Cmd(&resp, command, args...))
	if err != nil {
		return nil, err
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
