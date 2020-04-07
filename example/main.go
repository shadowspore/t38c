package main

import (
	"log"

	"github.com/k0kubun/pp"
	geojson "github.com/paulmach/go.geojson"
	t38c "github.com/zerobounty/tile38-client"
)

func main() {
	fencer, err := t38c.NewFence(t38c.NewRedisFencer("localhost:9851"))
	if err != nil {
		panic(err)
	}

	ch, err := fencer.FenceRoamIDs("people", "people", "*", 10000,
		nil,
	)
	if err != nil {
		panic(err)
	}

	for event := range ch {
		pp.Println(event)
	}

	return
	tile38, err := t38c.New(t38c.NewRadixPool("localhost:9851", 5), t38c.Debug)
	if err != nil {
		log.Fatal(err)
	}

	err = tile38.Set("keke", "popo", t38c.SetString("abce"))
	if err != nil {
		log.Fatal(err)
	}

	res, err := tile38.SearchCount("keke")
	if err != nil {
		log.Fatal(err)
	}
	pp.Println(res)
	return

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

	ob, err := tile38.GetPoint("fleet", "truck1", false)
	if err != nil {
		log.Fatal(err)
	}
	pp.Println(ob)

	tile38.Set("fleet", "truck2", t38c.SetGeometry(feat.Geometry))
	obe, err := tile38.Get("fleet", "truck2", false)
	if err != nil {
		log.Fatal(err)
	}
	pp.Println(obe.Object)

	pts, err := tile38.WithinPoints("fleet", t38c.AreaCircle(2, 2, 99999))
	if err != nil {
		log.Fatal(err)
	}

	pp.Println(pts)

	hashes, err := tile38.NearbyBounds("fleet", t38c.NearbyPoint(0, 0, 999999))
	if err != nil {
		log.Fatal(err)
	}

	pp.Println(hashes)
}
