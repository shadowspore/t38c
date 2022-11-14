//go:build ignore
// +build ignore

package main

import (
	"context"
	"log"

	geojson "github.com/paulmach/go.geojson"
	"github.com/xjem/t38c"
)

func main() {
	tile38, err := t38c.New(t38c.Config{
		Address: "localhost:9851",
		Debug:   true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer tile38.Close()

	tile38.Keys.Set("fleet", "truck1").Point(33.5123, -112.2693).Do(context.TODO())
	tile38.Keys.Set("fleet", "truck1").PointZ(33.5123, -112.2693, 225).Do(context.TODO())
	tile38.Keys.Set("fleet", "truck1").Bounds(30, -110, 40, -100).Do(context.TODO())
	tile38.Keys.Set("fleet", "truck1").Hash("9tbnthxzr").Do(context.TODO())

	polygon := geojson.NewPolygonGeometry([][][]float64{
		{
			{0, 0},
			{10, 10},
			{10, 0},
			{0, 0},
		},
	})
	tile38.Keys.Set("city", "tempe").Geometry(polygon).Do(context.TODO())
}
