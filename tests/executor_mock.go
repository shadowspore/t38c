package tests

import (
	"fmt"
	"sync"

	t38c "github.com/zerobounty/tile38-client"
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
	return func() (t38c.Executor, error) {
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
