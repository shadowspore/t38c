package main

import (
	"github.com/k0kubun/pp"

	t38c "github.com/zerobounty/tile38-client"
	"github.com/zerobounty/tile38-client/geofence"
)

func main() {
	fencer, err := geofence.NewRedisFencer("localhost:9851", true)
	if err != nil {
		panic(err)
	}

	req := geofence.NewFenceReq("WITHIN", "fleet", t38c.AreaCircle(0, 0, 1000)).Actions(geofence.Inside)
	ch, err := geofence.FenceObject(fencer, req)
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
