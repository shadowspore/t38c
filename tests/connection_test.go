package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	t38c "github.com/zerobounty/tile38-client"
)

func TestPing(t *testing.T) {
	mock := NewMockExecutor()

	tile38, err := t38c.New(mock.DialFunc(), t38c.Debug)
	assert.Nil(t, err)

	err = tile38.Ping()
	assert.Nil(t, err)
}
