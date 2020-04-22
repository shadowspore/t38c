package t38c

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/mediocregopher/radix/v3"
	"github.com/mediocregopher/radix/v3/resp/resp2"
)

var _ Executor = (*RadixPoolExecutor)(nil)

// RadixPoolExecutor struct
type RadixPoolExecutor struct {
	addr     string
	password *string
	pool     *radix.Pool
}

// NewRadixPool return radix pool dialer with provided pool size.
func NewRadixPool(addr string, size int) ExecutorDialer {
	return func(password *string) (Executor, error) {
		connFn := poolConnFn(password)
		pool, err := radix.NewPool("tcp", addr, size,
			radix.PoolConnFunc(connFn),
		)
		if err != nil {
			return nil, err
		}

		return &RadixPoolExecutor{
			addr:     addr,
			password: password,
			pool:     pool,
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

// ExecuteStream used for commands with streaming response.
// Creates a new connection for each stream.
func (rad *RadixPoolExecutor) ExecuteStream(ctx context.Context, command string, args ...string) (chan []byte, error) {
	conn, err := radix.Dial("tcp", rad.addr,
		radix.DialConnectTimeout(time.Second*10),
		radix.DialReadTimeout(0),
	)
	if err != nil {
		return nil, err
	}

	if err := radixPrepareConn(conn, rad.password); err != nil {
		conn.Close()
		return nil, err
	}

	var resp []byte
	if err := conn.Do(radix.Cmd(&resp, command, args...)); err != nil {
		conn.Close()
		return nil, err
	}

	if err := checkResponseErr(resp); err != nil {
		conn.Close()
		return nil, err
	}

	ch := make(chan []byte, 10)
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

func poolConnFn(password *string) radix.ConnFunc {
	return radix.ConnFunc(func(net, addr string) (radix.Conn, error) {
		conn, err := radix.Dial(net, addr,
			radix.DialConnectTimeout(time.Second*10),
		)
		if err != nil {
			return nil, err
		}

		if err := radixPrepareConn(conn, password); err != nil {
			conn.Close()
			return nil, err
		}

		return conn, nil
	})
}

func radixPrepareConn(conn radix.Conn, password *string) error {
	if password != nil {
		if err := conn.Do(radix.Cmd(nil, "AUTH", *password)); err != nil {
			return err
		}
	}

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
