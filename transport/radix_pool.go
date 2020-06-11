package transport

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/mediocregopher/radix/v3"
	"github.com/mediocregopher/radix/v3/resp/resp2"
	"github.com/tidwall/gjson"
)

// ErrClosedClient error
var ErrClosedClient = errors.New("closed client")

// RadixPool struct
type RadixPool struct {
	addr     string
	password *string
	pool     *radix.Pool
	wg       sync.WaitGroup
	closed   uint32
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
	if rad.isClosed() {
		return nil, ErrClosedClient
	}

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
	if rad.isClosed() {
		return ErrClosedClient
	}

	rad.wg.Add(1)
	defer rad.wg.Done()
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
		var response resp2.BulkStringBytes
		if err := conn.Decode(&response); err != nil {
			forced := atomic.LoadInt32(&forcedDisconnect) == 1
			if forced {
				return nil
			}

			return fmt.Errorf("resp decode: %v", err)
		}

		if err := handler(response.B); err != nil {
			return err
		}
	}

	return nil
}

func (rad *RadixPool) isClosed() bool {
	return atomic.LoadUint32(&rad.closed) == 1
}

// Close closes all connections in the pool and rejects future execution calls.
// Blocks until all streams are closed.
func (rad *RadixPool) Close() error {
	atomic.StoreUint32(&rad.closed, 1)
	err := rad.pool.Close()
	rad.pool = nil
	rad.wg.Wait()
	return err
}
