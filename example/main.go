package main

import (
	"github.com/k0kubun/pp"
	t38c "github.com/lostpeer/tile38-client"
)

func main() {
	fencer, err := t38c.NewRedisFencer("localhost:9851", true)
	if err != nil {
		panic(err)
	}

	req := t38c.NewFenceReq("WITHIN", "fleet", t38c.AreaCircle(0, 0, 1000)).Actions(t38c.Inside)
	ch, err := t38c.FenceObject(fencer, req)
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
