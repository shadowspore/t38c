// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"

	t38c "github.com/b3q/tile38-client"
)

func main() {
	tile38, err := t38c.New("localhost:9851", t38c.Debug, t38c.SetPoolSize(10))
	if err != nil {
		log.Fatal(err)
	}
	defer tile38.Close()

	// add a couple of points named 'truck1' and 'truck2' to a collection named 'fleet'.
	if err := tile38.Keys.Set("fleet", "truck1").Point(33.5123, -112.2693).Do(); err != nil {
		log.Fatal(err)
	}

	if err := tile38.Keys.Set("fleet", "truck2").Point(33.4626, -112.1695).Do(); err != nil {
		log.Fatal(err)
	}

	// search the 'fleet' collection.
	// returns both trucks in 'fleet'
	scanRes, err := tile38.Search.Scan("fleet").Do()
	if err != nil {
		log.Fatal(err)
	}
	printJSON("scan", scanRes)

	// search 6 kilometers around a point. returns one truck.
	nearbyRes, err := tile38.Search.Nearby("fleet", 33.462, -112.268, 6000).Do()
	if err != nil {
		log.Fatal(err)
	}
	printJSON("nearby", nearbyRes)

	// key value operations
	// returns 'truck1'
	truck1, err := tile38.Keys.Get("fleet", "truck1", false)
	if err != nil {
		log.Fatal(err)
	}
	printJSON("get truck1", truck1)

	// deletes 'truck2'
	if err := tile38.Keys.Del("fleet", "truck2"); err != nil {
		log.Fatal(err)
	}

	// removes all
	if err := tile38.Keys.Drop("fleet"); err != nil {
		log.Fatal(err)
	}
}

func printJSON(msg string, data interface{}) {
	b, _ := json.Marshal(data)
	fmt.Printf("%s: %s\n", msg, b)
}
