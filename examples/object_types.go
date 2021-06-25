// +build ignore

package main

import (
	"log"

	"github.com/paulmach/orb"
	geojson "github.com/paulmach/orb/geojson"
	"github.com/xjem/t38c"
)

func main() {
	tile38, err := t38c.New("localhost:9851", t38c.Debug)
	if err != nil {
		log.Fatal(err)
	}
	defer tile38.Close()

	tile38.Keys.Set("fleet", "truck1").Point(33.5123, -112.2693).Do()       // nolint:errcheck
	tile38.Keys.Set("fleet", "truck1").PointZ(33.5123, -112.2693, 225).Do() // nolint:errcheck
	tile38.Keys.Set("fleet", "truck1").Bounds(30, -110, 40, -100).Do()      // nolint:errcheck
	tile38.Keys.Set("fleet", "truck1").Hash("9tbnthxzr").Do()               // nolint:errcheck

	polygon := geojson.NewGeometry(orb.Polygon{
		{
			{0, 0},
			{10, 10},
			{10, 0},
			{0, 0},
		},
	})
	tile38.Keys.Set("city", "tempe").Geometry(polygon).Do() // nolint:errcheck
}
