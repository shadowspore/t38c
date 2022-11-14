package e2e

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	geojson "github.com/paulmach/go.geojson"
	"github.com/stretchr/testify/require"
	"github.com/xjem/t38c"
)

func testGeofence(t *testing.T, client *t38c.Client) {
	go func() {
		time.Sleep(time.Second)

		err := client.Keys.Set("geofence-test", "1").Point(33, -112).Do(context.Background())
		require.NoError(t, err)
		err = client.Keys.Set("geofence-test", "1").Point(40, -150).Do(context.Background())
		require.NoError(t, err)
	}()

	errOk := errors.New("success")
	n := 0
	handler := t38c.EventHandlerFunc(func(event *t38c.GeofenceEvent) error {
		b, err := json.Marshal(event)
		if err != nil {
			return err
		}

		fmt.Printf("event: %s\n", b)
		switch n {
		case 0:
			require.Equal(t, event.Command, "set")
			require.Equal(t, event.Detect, "enter")
			require.Equal(t, event.Key, "geofence-test")
			require.Equal(t, event.ID, "1")
			require.Equal(t, event.Object, &t38c.Object{
				Geometry: &geojson.Geometry{
					Type:  geojson.GeometryPoint,
					Point: []float64{-112, 33},
				},
			})
		case 1:
			require.Equal(t, event.Command, "set")
			require.Equal(t, event.Detect, "exit")
			require.Equal(t, event.Key, "geofence-test")
			require.Equal(t, event.ID, "1")
			require.Equal(t, event.Object, &t38c.Object{
				Geometry: &geojson.Geometry{
					Type:  geojson.GeometryPoint,
					Point: []float64{-150, 40},
				},
			})
			return errOk
		}

		n++
		return nil
	})

	err := client.Geofence.Nearby("geofence-test", 33, -112, 10000).
		Actions(t38c.Enter, t38c.Exit).
		Do(context.Background(), handler)
	require.EqualError(t, err, errOk.Error())
}
