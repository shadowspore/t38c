package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	t38c "github.com/lostpeer/tile38-client"
)

func TestPing(t *testing.T) {
	conn := NewMockedConn()
	tile38, err := t38c.NewWithConn(conn)
	assert.Nil(t, err)

	err = tile38.Ping()
	assert.Nil(t, err)
}
