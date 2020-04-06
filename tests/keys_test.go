package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	t38c "github.com/lostpeer/tile38-client"
)

func TestBounds(t *testing.T) {
	pool, err := NewMocker().
		Mock(
			`BOUNDS test`, `{"ok":true,"bounds":{"type":"Polygon","coordinates":[[[1,1],[2,1],[2,2],[1,2],[1,1]]]},"elapsed":"19.52µs"}`,
		).GetPool()
	assert.Nil(t, err)

	tile38, err := t38c.New(t38c.ClientOptions{
		Pool: pool,
		Debug: true,
	})
	assert.Nil(t, err)

	resp, err := tile38.Bounds("test")
	assert.Nil(t, err)

	expected := [][][]float64{
		[][]float64{
			[]float64{1, 1},
			[]float64{2, 1},
			[]float64{2, 2},
			[]float64{1, 2},
			[]float64{1, 1},
		},
	}
	assert.Equal(t, expected, resp)
}
