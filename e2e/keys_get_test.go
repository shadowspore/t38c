package e2e

import (
	"context"
	"testing"
	"time"

	geojson "github.com/paulmach/go.geojson"
	"github.com/stretchr/testify/require"
	"github.com/sythang/t38c"
)

func testKeys(t *testing.T, client *t38c.Client) {
	testKeysGet(t, client)

	// Expiration tests.
	err := client.Keys.Set("foo", "baz").Bounds(0, 0, 20, 20).Field("age", 20).Expiration(2).Do(context.Background())
	require.NoError(t, err)

	resp, err := client.Keys.Get("foo", "baz").WithFields().Object(context.Background())
	require.NoError(t, err)

	require.Equal(t, geojson.NewPolygonGeometry([][][]float64{{
		{0, 0},
		{20, 0},
		{20, 20},
		{0, 20},
		{0, 0},
	}}), resp.Object.Geometry)

	time.Sleep(time.Second * 3)

	_, err = client.Keys.Get("foo", "baz").Object(context.Background())
	require.Error(t, err)
}

func testKeysGet(t *testing.T, client *t38c.Client) {
	require.NoError(t, client.Keys.Set("foo", "bar").Point(1, 2).Do(context.Background()))
	// Check object.
	{
		resp, err := client.Keys.Get("foo", "bar").Object(context.Background())
		require.NoError(t, err)
		require.Equal(t, geojson.NewPointGeometry([]float64{2, 1}), resp.Object.Geometry)
	}

	// Check point.
	{
		resp, err := client.Keys.Get("foo", "bar").Point(context.Background())
		require.NoError(t, err)
		require.Equal(t, t38c.Point{
			Lat: 1,
			Lon: 2,
		}, resp.Point)
	}

	// Check bounds.
	{
		resp, err := client.Keys.Get("foo", "bar").Bounds(context.Background())
		require.NoError(t, err)
		require.Equal(t, t38c.Bounds{
			SW: t38c.Point{Lat: 1, Lon: 2},
			NE: t38c.Point{Lat: 1, Lon: 2},
		}, resp.Bounds)
	}

	// Check hash.
	{
		resp, err := client.Keys.Get("foo", "bar").Hash(context.Background(), 5)
		require.NoError(t, err)
		require.Equal(t, "s01mt", resp.Hash)
	}
}
