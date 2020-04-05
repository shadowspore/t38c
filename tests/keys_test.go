package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	t38c "github.com/lostpeer/tile38-client"
)

func TestBounds(t *testing.T) {
	conn := NewMockedConn()
	conn.Mock("BOUNDS test", []byte(`
		{"ok":true,"bounds":{"type":"Polygon","coordinates":[[[1,1],[2,1],[2,2],[1,2],[1,1]]]},"elapsed":"19.52µs"}
	`))

	tile38, err := t38c.NewWithConn(conn)
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
