package geofence

import (
	"fmt"
	"time"

	"log"

	"github.com/mediocregopher/radix/v3"
	"github.com/mediocregopher/radix/v3/resp/resp2"
	"github.com/tidwall/gjson"
	t38c "github.com/lostpeer/tile38-client"
)

// RadixFencer struct
type RadixFencer struct {
	addr string
}

// NewRadixFencer ...
func NewRadixFencer(addr string) ExecutorDialer {
	return func() (Executor, error) {
		conn, err := radix.Dial("tcp", addr,
			radix.DialConnectTimeout(time.Second*10),
		)
		if err != nil {
			return nil, err
		}

		conn.Close()
		return &RadixFencer{addr: addr}, nil
	}
}

// Fence ...
func (fencer *RadixFencer) Fence(command string, args ...string) (ch chan []byte, err error) {
	conn, err := radix.Dial("tcp", fencer.addr,
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
		if err := t38c.RadixJSONifyConn(conn); err != nil {
			return nil, err
		}

		var resp []byte
		if err := conn.Do(radix.Cmd(&resp, command, args...)); err != nil {
			return nil, err
		}

		if !gjson.GetBytes(resp, "ok").Bool() {
			return nil, fmt.Errorf(gjson.GetBytes(resp, "err").String())
		}

		if !gjson.GetBytes(resp, "live").Bool() {
			return nil, fmt.Errorf("not live: %s", resp)
		}
	}

	ch = make(chan []byte, 10)
	go func() {
		defer func() {
			close(ch)
			conn.Close()
		}()
		for {
			resp := &resp2.BulkStringBytes{}
			if err := conn.Decode(resp); err != nil {
				log.Printf("resp decode: %v\n", err)
				break
			}

			ch <- resp.B
		}

	}()

	return ch, nil
}
