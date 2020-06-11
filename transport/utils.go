package transport

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mediocregopher/radix/v3"
)

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
