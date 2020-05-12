package transport

import (
	"context"
	"fmt"
	"sync"
)

// Mock struct
type Mock struct {
	mut   sync.Mutex
	mocks map[string]string
}

// NewMock ...
func NewMock() *Mock {
	m := &Mock{
		mocks: make(map[string]string),
	}

	m.Mock(`PING`, `{"ok": true, "ping": "pong"}`)
	m.Mock(`OUTPUT json`, `{"ok": true}`)

	return m
}

// Mock ...
func (m *Mock) Mock(request, response string) {
	m.mut.Lock()
	defer m.mut.Unlock()
	m.mocks[request] = response
}

// Execute ...
func (m *Mock) Execute(command string, args ...string) ([]byte, error) {
	m.mut.Lock()
	defer m.mut.Unlock()

	cmd := command
	for _, arg := range args {
		cmd += " " + arg
	}

	resp, found := m.mocks[cmd]
	if !found {
		return nil, fmt.Errorf("request '%s' is not set", cmd)
	}

	return []byte(resp), nil
}

// ExecuteStream ...
func (m *Mock) ExecuteStream(ctx context.Context, handler func([]byte) error, command string, args ...string) error {
	return fmt.Errorf("not implemented")
}
