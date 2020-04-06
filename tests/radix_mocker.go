package tests

import (
	"fmt"
	"strings"
	"sync"

	"github.com/mediocregopher/radix/v3"
)

// Mocker struct
type Mocker struct {
	mu    sync.Mutex
	mocks map[string]string
}

// NewMocker ...
func NewMocker() *Mocker {
	m := &Mocker{
		mocks: make(map[string]string),
	}

	m.Mock(`PING`, `{"ping": "pong"}`)
	m.Mock(`OUTPUT json`, ``)

	return m
}

// Mock ...
func (m *Mocker) Mock(request string, response string) *Mocker {
	m.mocks[request] = response
	return m
}

// GetPool ...
func (m *Mocker) GetPool() (*radix.Pool, error) {
	conn := radix.Stub("tcp", "localhost:9851", func(args []string) interface{} {
		req := strings.Join(args, " ")
		resp, found := m.mocks[req]
		if !found {
			return fmt.Errorf("request '%s' not specified", req)
		}

		return []byte(resp)
	})

	customConnFunc := func(network, addr string) (radix.Conn, error) {
		return conn, nil
	}

	return radix.NewPool("tcp", "localhost:9851", 1, radix.PoolConnFunc(customConnFunc))

}
