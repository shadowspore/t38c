package main

import (
	"log"

	"github.com/k0kubun/pp"
	t38c "github.com/lostpeer/tile38-client"
)

func main() {
	tile38, err := t38c.New(t38c.NewRadixPool("localhost:9851", 5), t38c.Debug)
	if err != nil {
		log.Fatal(err)
	}

	req := t38c.Nearby("fleet", t38c.NearbyPoint(0, 0, 999999)).
		WithOptions(t38c.Distance()).Format(t38c.OutputPoints)

	resp, err := tile38.Search(req)
	if err != nil {
		log.Fatal(err)
	}

	pp.Println(resp)
}
