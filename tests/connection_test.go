package tests

import (
	"testing"

	t38c "github.com/b3q/tile38-client"
	"github.com/b3q/tile38-client/transport"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	tile38, err := t38c.NewWithExecutor(transport.NewMock(), true)
	assert.Nil(t, err)

	err = tile38.Ping()
	assert.Nil(t, err)
}
