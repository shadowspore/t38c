package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/mediocregopher/radix/v3"
	"github.com/mediocregopher/radix/v3/resp/resp2"
	"github.com/tidwall/gjson"
)

// RadixPool struct
type RadixPool struct {
	addr     string
	password *string
	pool     *radix.Pool
}

// NewRadixPool return radix pool dialer with provided pool size.
func NewRadixPool(addr string, size int, password *string) (*RadixPool, error) {
	connFn := poolConnFn(password)
	pool, err := radix.NewPool("tcp", addr, size,
		radix.PoolConnFunc(connFn),
	)
	if err != nil {
		return nil, err
	}

	return &RadixPool{
		addr:     addr,
		password: password,
		pool:     pool,
	}, nil
}

// Execute command
func (rad *RadixPool) Execute(command string, args ...string) ([]byte, error) {
	var resp []byte
	err := rad.pool.Do(radix.Cmd(&resp, command, args...))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ExecuteStream used for commands with streaming response.
// Creates a new connection for each stream.
func (rad *RadixPool) ExecuteStream(ctx context.Context, handler func([]byte) error, command string, args ...string) error {
	conn, err := radix.Dial("tcp", rad.addr,
		radix.DialConnectTimeout(time.Second*10),
		radix.DialReadTimeout(0),
	)
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := radixPrepareConn(conn, rad.password); err != nil {
		return err
	}

	var resp []byte
	if err := conn.Do(radix.Cmd(&resp, command, args...)); err != nil {
		return err
	}

	if !gjson.GetBytes(resp, "ok").Bool() {
		return fmt.Errorf(gjson.GetBytes(resp, "err").String())
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var forcedDisconnect int32
	go func() {
		<-ctx.Done()
		atomic.StoreInt32(&forcedDisconnect, 1)
		conn.Close()
	}()

	for atomic.LoadInt32(&forcedDisconnect) == 0 {
		resp := &resp2.BulkStringBytes{}
		if err := conn.Decode(resp); err != nil {
			forced := atomic.LoadInt32(&forcedDisconnect) == 1
			if forced {
				return nil
			}

			return fmt.Errorf("resp decode: %v", err)
		}

		if err := handler(resp.B); err != nil {
			return err
		}
	}

	return nil
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
