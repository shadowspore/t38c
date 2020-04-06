package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	t38c "github.com/lostpeer/tile38-client"
)

func TestPing(t *testing.T) {
	pool, err := NewMocker().GetPool()
	assert.Nil(t, err)

	tile38, err := t38c.NewWithPool(pool, t38c.Debug())
	assert.Nil(t, err)

	err = tile38.Ping()
	assert.Nil(t, err)
}
