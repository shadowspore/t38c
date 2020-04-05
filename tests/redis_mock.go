package tests

import (
	"fmt"
	"sync"

	"github.com/garyburd/redigo/redis"
)

var _ = redis.Conn(&RedisMockedConn{})

// RedisMockedConn struct
type RedisMockedConn struct {
	mtx         *sync.Mutex
	mockedResps map[string][]byte
}

// NewMockedConn ...
func NewMockedConn() *RedisMockedConn {
	conn := &RedisMockedConn{
		mtx:         &sync.Mutex{},
		mockedResps: make(map[string][]byte, 0),
	}
	conn.Mock("PING", []byte(`{"ping": "pong"}`))
	conn.Mock("OUTPUT json", nil)

	return conn
}

// Mock ...
func (conn *RedisMockedConn) Mock(request string, response []byte) {
	conn.mtx.Lock()
	defer conn.mtx.Unlock()
	conn.mockedResps[request] = response
}

// Do ...
func (conn *RedisMockedConn) Do(command string, args ...interface{}) (interface{}, error) {
	conn.mtx.Lock()
	defer conn.mtx.Unlock()

	req := requestToString(command, args)
	resp, found := conn.mockedResps[req]
	if !found {
		return nil, fmt.Errorf("response for request '%s' not specified", req)
	}

	return resp, nil
}

// Send ...
func (conn *RedisMockedConn) Send(command string, args ...interface{}) error {
	return fmt.Errorf("not implemented")
}

// Err ...
func (conn *RedisMockedConn) Err() error {
	return fmt.Errorf("not implemented")
}

// Close ...
func (conn *RedisMockedConn) Close() error {
	return fmt.Errorf("not implemented")
}

// Flush ...
func (conn *RedisMockedConn) Flush() error {
	return fmt.Errorf("not implemented")
}

// Receive ...
func (conn *RedisMockedConn) Receive() (interface{}, error) {
	return nil, fmt.Errorf("not implemented")
}
