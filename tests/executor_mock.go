package tests

import (
	"context"
	"fmt"
	"sync"

	t38c "github.com/powercake/tile38-client"
)

var _ t38c.Executor = (*MockExecutor)(nil)

// MockExecutor struct
type MockExecutor struct {
	mut   sync.Mutex
	mocks map[string]string
}

// NewMockExecutor ...
func NewMockExecutor() *MockExecutor {
	m := &MockExecutor{
		mocks: make(map[string]string),
	}

	m.Mock(`PING`, `{"ok": true, "ping": "pong"}`)
	m.Mock(`OUTPUT json`, `{"ok": true}`)

	return m
}

// DialFunc ...
func (m *MockExecutor) DialFunc() t38c.ExecutorDialer {
	return func(password *string) (t38c.Executor, error) {
		return m, nil
	}
}

// Mock ...
func (m *MockExecutor) Mock(request, response string) {
	m.mut.Lock()
	defer m.mut.Unlock()
	m.mocks[request] = response
}

// Execute ...
func (m *MockExecutor) Execute(command string, args ...string) ([]byte, error) {
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
func (m *MockExecutor) ExecuteStream(ctx context.Context, command string, args ...string) (chan []byte, error) {
	return nil, fmt.Errorf("not implemented")
}
