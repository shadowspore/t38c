package t38c

import (
	"encoding/json"
	"fmt"
	"time"

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
		pool, err := radix.NewPool("tcp", addr, size, radix.PoolConnFunc(func(net, addr string) (conn radix.Conn, err error) {
			conn, err = radix.Dial(net, addr,
				radix.DialConnectTimeout(time.Second*10),
			)
			if err != nil {
				return
			}

			err = RadixJSONifyConn(conn)
			if err != nil {
				conn.Close()
			}
			return
		}))
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

// RadixJSONifyConn ...
func RadixJSONifyConn(conn radix.Conn) (err error) {
	var b []byte
	if err := conn.Do(radix.Cmd(&b, "OUTPUT", "json")); err != nil {
		return err
	}

	var resp struct {
		Ok bool `json:"ok"`
	}
	if err := json.Unmarshal(b, &resp); err != nil {
		return err
	}

	if !resp.Ok {
		return fmt.Errorf("bad response: %s", b)
	}

	return nil
}
