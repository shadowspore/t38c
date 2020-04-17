package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"

	t38c "github.com/lostpeer/tile38-client"
)

func printJSON(msg string, data interface{}) {
	b, _ := json.Marshal(data)
	fmt.Printf("%s: %s\n", msg, b)
}

func main() {
	tile38, err := t38c.New("localhost:9851", true)
	if err != nil {
		log.Fatal(err)
	}

	results, err := tile38.Search(
		t38c.Nearby("fleet", 33.462, -112.268, 6000,
			t38c.Where("speed", 70, math.MaxInt32),
			t38c.Match("truck*"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	printJSON("nearby:", results)
}
