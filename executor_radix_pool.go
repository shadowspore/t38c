package t38c

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/mediocregopher/radix/v3"
	"github.com/mediocregopher/radix/v3/resp/resp2"
	"github.com/tidwall/gjson"
)

var _ Executor = (*RadixPoolExecutor)(nil)

// RadixPoolExecutor struct
type RadixPoolExecutor struct {
	addr string
	pool *radix.Pool
}

// NewRadixPool ...
func NewRadixPool(addr string, size int) ExecutorDialer {
	return func() (Executor, error) {
		pool, err := radix.NewPool("tcp", addr, size,
			radix.PoolConnFunc(poolConnFn),
		)
		if err != nil {
			return nil, err
		}

		return &RadixPoolExecutor{
			addr: addr,
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

// ExecuteStream ...
func (rad *RadixPoolExecutor) ExecuteStream(ctx context.Context, command string, args ...string) (ch chan []byte, err error) {
	conn, err := radix.Dial("tcp", rad.addr,
		radix.DialConnectTimeout(time.Second*10),
		radix.DialReadTimeout(0),
	)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			conn.Close()
		}
	}()

	{
		if err := radixJSONifyConn(conn); err != nil {
			return nil, err
		}

		var resp []byte
		if err := conn.Do(radix.Cmd(&resp, command, args...)); err != nil {
			return nil, err
		}

		if err := checkResponseErr(resp); err != nil {
			return nil, err
		}

		if !gjson.GetBytes(resp, "live").Bool() {
			return nil, fmt.Errorf("bad response: %s", resp)
		}
	}

	ch = make(chan []byte, 10)
	go func() {
		defer func() {
			conn.Close()
			close(ch)
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				resp := &resp2.BulkStringBytes{}
				if err := conn.Decode(resp); err != nil {
					log.Printf("resp decode: %v\n", err)
					return
				}

				ch <- resp.B
			}
		}
	}()

	return ch, nil
}

func poolConnFn(net, addr string) (conn radix.Conn, err error) {
	conn, err = radix.Dial(net, addr,
		radix.DialConnectTimeout(time.Second*10),
	)
	if err != nil {
		return
	}

	err = radixJSONifyConn(conn)
	if err != nil {
		conn.Close()
	}
	return
}

// RadixJSONifyConn ...
func radixJSONifyConn(conn radix.Conn) (err error) {
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
