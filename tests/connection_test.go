package tests

import (
	"testing"

	t38c "github.com/powercake/tile38-client"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	mock := NewMockExecutor()

	tile38, err := t38c.NewWithDialer(mock.DialFunc(), t38c.Debug())
	assert.Nil(t, err)

	err = tile38.Ping()
	assert.Nil(t, err)
}
