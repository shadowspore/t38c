package main

import (
	"log"

	"github.com/k0kubun/pp"

	t38c "github.com/zerobounty/tile38-client"
	"github.com/zerobounty/tile38-client/geofence"
)

func main() {
	geo, err := geofence.New(geofence.NewRedisFencer("localhost:9851"), true)
	if err != nil {
		log.Fatal(err)
	}

	ch, err := geo.FenceIntersectsObjects(
		geofence.NewFenceReq("fleet", t38c.AreaCircle(0, 0, 1000)).
			Actions(geofence.Enter, geofence.Exit).
			WithOptions(t38c.Where("speed", 10, 20)),
	)
	if err != nil {
		log.Fatal(err)
	}
	for event := range ch {
		pp.Println(event)
	}
	// tile38, err := t38c.New(t38c.NewRadixPool("localhost:9851", 1), t38c.Debug)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// res, err := tile38.IntersectsHashes("fleet", t38c.AreaCircle(0, 0, 999999), 5)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// pp.Println(res)
}
