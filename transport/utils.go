package transport

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mediocregopher/radix/v4"
)

func poolConnFn(password *string) func(ctx context.Context, net, addr string) (radix.Client, error) {
	return func(ctx context.Context, net, addr string) (radix.Client, error) {
		conn, err := radix.Dial(ctx, net, addr)
		if err != nil {
			return nil, err
		}

		if err := radixPrepareConn(ctx, conn, password); err != nil {
			_ = conn.Close()
			return nil, err
		}

		return conn, nil
	}
}

func radixPrepareConn(ctx context.Context, conn radix.Conn, password *string) error {
	if password != nil {
		if err := conn.Do(ctx, radix.Cmd(nil, "AUTH", *password)); err != nil {
			return err
		}
	}

	var b []byte
	if err := conn.Do(ctx, radix.Cmd(&b, "OUTPUT", "json")); err != nil {
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
