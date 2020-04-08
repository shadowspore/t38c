package main

import (
	"log"

	"github.com/k0kubun/pp"
	t38c "github.com/zerobounty/tile38-client"
)

func main() {
	tile38, err := t38c.New(t38c.NewRadixPool("localhost:9851", 5), t38c.Debug)
	if err != nil {
		log.Fatal(err)
	}

	res, err := tile38.Nearby("fleet", t38c.NearbyPoint(0, 0, 9999999))
	if err != nil {
		log.Fatal(err)
	}

	pp.Println(res)

	res1, err := tile38.Get("fleet", "truck1", true)
	if err != nil {
		log.Fatal(err)
	}

	pp.Println(res1)
}
