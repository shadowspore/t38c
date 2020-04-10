package geofence

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/tidwall/gjson"
)

var _ Executor = (*RedisFencer)(nil)

// RedisFencer struct
type RedisFencer struct {
	addr string
}

// NewRedisFencer ...
func NewRedisFencer(addr string) ExecutorDialer {
	return func() (Executor, error) {
		conn, err := dialRedis(addr)
		if err != nil {
			return nil, err
		}
		defer conn.Close()

		fencer := &RedisFencer{
			addr: addr,
		}

		return fencer, nil
	}
}

// Fence ...
func (fencer *RedisFencer) Fence(command string, args ...string) (ch chan []byte, err error) {
	conn, err := dialRedis(fencer.addr)
	if err != nil {
		return nil, err
	}

	var ifaceArgs []interface{}
	for _, arg := range args {
		ifaceArgs = append(ifaceArgs, arg)
	}

	resp, err := conn.Do(command, ifaceArgs...)
	if err != nil {
		conn.Close()
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

	ch = make(chan []byte, 10)
	go func() {
		defer func() {
			close(ch)
			conn.Close()
		}()

		for {
			resp, err = conn.Receive()
			if err != nil {
				if err == io.EOF {
					break
				}

				log.Printf("receive: %v\n", err)
				break
			}

			body, ok = resp.([]byte)
			if !ok {
				log.Printf("bad response type: %v", resp)
				break
			}

			ch <- body
		}
	}()

	return ch, nil
}

func dialRedis(addr string) (redis.Conn, error) {
	conn, err := redis.Dial("tcp", addr, redis.DialConnectTimeout(time.Second*10))
	if err != nil {
		return nil, err
	}

	if _, err := conn.Do("OUTPUT", "json"); err != nil {
		return nil, err
	}

	return conn, nil
}
