package e2e

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/paulmach/orb"
	geojson "github.com/paulmach/orb/geojson"
	"github.com/stretchr/testify/require"
	"github.com/xjem/t38c"
)

func TestE2E(t *testing.T) {
	skipE2E(t)

	client, err := t38c.New(os.Getenv("T38C_TEST_ADDR"))
	require.NoError(t, err)

	t.Run("TestKeys", testKeys(client))
	t.Run("TestWithin", testWithin(client))
}

func testKeys(client *t38c.Client) func(t *testing.T) {
	return func(t *testing.T) {
		require.NoError(t, client.Keys.Set("foo", "bar").Point(0, 0).Do())

		resp, err := client.Keys.Get("foo", "bar", false)
		require.NoError(t, err)

		require.Equal(t, orb.Point{0, 0}, resp.Object.Geometry)

		err = client.Keys.Set("foo", "baz").Bounds(0, 0, 20, 20).Field("age", 20).Expiration(2).Do()
		require.NoError(t, err)

		resp, err = client.Keys.Get("foo", "baz", true)
		require.NoError(t, err)

		require.Equal(t, orb.Polygon{{
			{0, 0},
			{20, 0},
			{20, 20},
			{0, 20},
			{0, 0},
		}}, resp.Object.Geometry)

		time.Sleep(time.Second * 3)

		_, err = client.Keys.Get("foo", "baz", false)
		require.Error(t, err)
	}
}

func testWithin(client *t38c.Client) func(t *testing.T) {
	return func(t *testing.T) {
		err := client.Keys.Set("points", "point-1").Point(1, 1).Do()
		require.NoError(t, err)

		geom := &geojson.Geometry{
			Coordinates: orb.Polygon{{
				{0, 0},
				{20, 0},
				{20, 20},
				{0, 20},
				{0, 0},
			}},
		}

		err = client.Keys.Set("areas", "area-1").Geometry(geom).Do()
		require.NoError(t, err)

		resp, err := client.Search.Within("points").
			Get("areas", "area-1").
			Format(t38c.FormatIDs).Do()
		require.NoError(t, err)

		require.Equal(t, []string{"point-1"}, resp.IDs)
	}

}

func skipE2E(tb testing.TB) {
	env := "T38C_TEST_E2E"

	tb.Helper()
	if ok, _ := strconv.ParseBool(os.Getenv(env)); !ok {
		tb.Skipf("Skipped. Set %s=1 to enable e2e test.", env)
	}
}
