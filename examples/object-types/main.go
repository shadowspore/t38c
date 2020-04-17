package main

import (
	"log"

	t38c "github.com/lostpeer/tile38-client"
	geojson "github.com/paulmach/go.geojson"
)

func main() {
	tile38, err := t38c.New("localhost:9851", true)
	if err != nil {
		log.Fatal(err)
	}

	tile38.Set("fleet", "truck1", t38c.SetPoint(33.5123, -112.2693))
	tile38.Set("fleet", "truck1", t38c.SetPointZ(33.5123, -112.2693, 225))
	tile38.Set("fleet", "truck1", t38c.SetBounds(30, -110, 40, -100))
	tile38.Set("fleet", "truck1", t38c.SetHash("9tbnthxzr"))

	polygon := geojson.NewPolygonGeometry([][][]float64{
		{
			{0, 0},
			{10, 10},
			{10, 0},
			{0, 0},
		},
	})
	tile38.Set("city", "tempe", t38c.SetGeometry(polygon))
}
