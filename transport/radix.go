package transport

import (
	"context"
	"fmt"

	"github.com/mediocregopher/radix/v4"
	"github.com/mediocregopher/radix/v4/resp/resp3"
	"github.com/tidwall/gjson"
)

// Radix is an t38c.Executor implementation based on 'mediocregopher/radix' library.
type Radix struct {
	addr     string
	password *string
	client   radix.Client

	ctx    context.Context
	cancel context.CancelFunc
}

// NewRadix returns radix transport implementation with provided pool size.
func NewRadix(addr string, size int, password *string) (*Radix, error) {
	client, err := (radix.PoolConfig{
		CustomPool: poolConnFn(password),
		Size:       size,
	}).New(context.Background(), "tcp", addr)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Radix{
		addr:     addr,
		password: password,
		client:   client,
		ctx:      ctx,
		cancel:   cancel,
	}, nil
}

// Execute executes command.
func (r *Radix) Execute(ctx context.Context, command string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var resp []byte
	if err := r.client.Do(ctx, radix.Cmd(&resp, command, args...)); err != nil {
		return nil, err
	}

	return resp, nil
}

// ExecuteStream used for commands with streaming response.
// Creates a new connection for each stream.
func (r *Radix) ExecuteStream(ctx context.Context, handler func([]byte) error, command string, args ...string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
		case <-r.ctx.Done():
			cancel()
		}
	}()

	conn, err := radix.Dial(ctx, "tcp", r.addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := radixPrepareConn(ctx, conn, r.password); err != nil {
		return err
	}

	var resp []byte
	if err := conn.Do(ctx, radix.Cmd(&resp, command, args...)); err != nil {
		return err
	}

	if !gjson.GetBytes(resp, "ok").Bool() {
		return fmt.Errorf(gjson.GetBytes(resp, "err").String())
	}

	for {
		var response resp3.BlobStringBytes
		if err := conn.EncodeDecode(ctx, nil, &response); err != nil {
			return fmt.Errorf("decode response: %v", err)
		}

		if err := handler(response.B); err != nil {
			return err
		}
	}
}

// Close closes the connections.
// All future execution calls will return an error.
func (r *Radix) Close() error {
	r.cancel()
	return r.client.Close()
}
