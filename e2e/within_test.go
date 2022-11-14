package e2e

import (
	"context"
	"testing"

	geojson "github.com/paulmach/go.geojson"
	"github.com/stretchr/testify/require"
	"github.com/xjem/t38c"
)

func testWithin(t *testing.T, client *t38c.Client) {
	err := client.Keys.Set("points", "point-1").Point(1, 1).Do(context.Background())
	require.NoError(t, err)

	geom := geojson.NewPolygonGeometry([][][]float64{{
		{0, 0},
		{20, 0},
		{20, 20},
		{0, 20},
		{0, 0},
	}})

	err = client.Keys.Set("areas", "area-1").Geometry(geom).Do(context.Background())
	require.NoError(t, err)

	resp, err := client.Search.Within("points").
		Get("areas", "area-1").
		Format(t38c.FormatIDs).Do(context.Background())
	require.NoError(t, err)

	require.Equal(t, []string{"point-1"}, resp.IDs)
}
