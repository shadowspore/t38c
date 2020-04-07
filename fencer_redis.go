package t38c

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/tidwall/gjson"
)

var _ FenceExecutor = (*RedisFencer)(nil)

type RedisFencer struct {
	conn redis.Conn
}

func NewRedisFencer(addr string) FenceExecutorDialer {
	return func() (FenceExecutor, error) {
		conn, err := redis.Dial("tcp", addr)
		if err != nil {
			return nil, err
		}

		if _, err := conn.Do("OUTPUT", "json"); err != nil {
			return nil, err
		}

		fencer := &RedisFencer{
			conn: conn,
		}

		return fencer, nil
	}
}

func (fencer *RedisFencer) Fence(command string, args ...string) (FenceChan, error) {
	var ifaceArgs []interface{}
	for _, arg := range args {
		ifaceArgs = append(ifaceArgs, arg)
	}

	resp, err := fencer.conn.Do(command, ifaceArgs...)
	if err != nil {
		return nil, err
	}

	body, ok := resp.([]byte)
	if !ok {
		return nil, fmt.Errorf("bad response type: %v", resp)
	}

	if !gjson.GetBytes(body, "ok").Bool() {
		return nil, fmt.Errorf(gjson.GetBytes(body, "err").String())
	}

	if !gjson.GetBytes(body, "live").Bool() {
		return nil, fmt.Errorf("not live: %v", resp)
	}

	ch := make(FenceChan, 10)
	go func() {
		for {
			resp, err = fencer.conn.Receive()
			if err != nil {
				ch <- FenceEvent{
					Err: err,
				}
				continue
			}

			body, ok = resp.([]byte)
			if !ok {
				ch <- FenceEvent{
					Err: fmt.Errorf("bad response type: %v", resp),
				}
				continue
			}

			ch <- FenceEvent{
				Data: body,
			}
		}
	}()

	return ch, nil
}
