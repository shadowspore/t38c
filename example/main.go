package main

import (
	"log"

	"github.com/k0kubun/pp"
	geojson "github.com/paulmach/go.geojson"
	t38c "github.com/zerobounty/tile38-client"
)

func main() {
	tile38, err := t38c.New("localhost:9851", t38c.Debug())
	if err != nil {
		log.Fatal(err)
	}

	tile38.Bounds("fleet")
	feat := geojson.NewPolygonFeature([][][]float64{
		[][]float64{
			[]float64{1, 1},
			[]float64{2, 1},
			[]float64{1, 2},
			[]float64{1, 1},
			[]float64{1, 1},
		},
	})

	feat.SetProperty("sadness", 999)
	if err := tile38.Set("fleet", "truck01", t38c.SetFeature(feat),
		t38c.SetField("speed", 10),
	); err != nil {
		log.Fatal(err)
	}

	obs, err := tile38.Intersects("fleet", t38c.AreaCircle(0, 0, 999999999))
	if err != nil {
		log.Fatal(err)
	}
	pp.Println(obs)

	ob, err := tile38.GetPoint("fleet", "truck1")
	if err != nil {
		log.Fatal(err)
	}
	pp.Println(ob)

	tile38.Set("fleet", "truck2", t38c.SetGeometry(feat.Geometry))
	obe, err := tile38.Get("fleet", "truck2", false)
	if err != nil {
		log.Fatal(err)
	}
	pp.Println(obe.GeoJSON)

	pts, err := tile38.WithinPoints("fleet", t38c.AreaCircle(2, 2, 99999))
	if err != nil {
		log.Fatal(err)
	}

	pp.Println(pts)
}
