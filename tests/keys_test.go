package tests

import (
	"testing"

	t38c "github.com/b3q/tile38-client"
	"github.com/b3q/tile38-client/transport"
	"github.com/stretchr/testify/assert"
)

func TestBounds(t *testing.T) {
	mock := transport.NewMock()
	mock.Mock(
		`BOUNDS test`,
		`{"ok":true,"bounds":{"type":"Polygon","coordinates":[[[1,1],[2,1],[2,2],[1,2],[1,1]]]},"elapsed":"19.52µs"}`,
	)

	tile38, err := t38c.NewWithExecutor(mock, true)
	assert.Nil(t, err)

	resp, err := tile38.Bounds("test")
	assert.Nil(t, err)

	expected := [][][]float64{
		{
			[]float64{1, 1},
			[]float64{2, 1},
			[]float64{2, 2},
			[]float64{1, 2},
			[]float64{1, 1},
		},
	}
	assert.Equal(t, expected, resp)
}
