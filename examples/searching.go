// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"

	t38c "github.com/powercake/tile38-client"
)

func main() {
	tile38, err := t38c.New("localhost:9851", t38c.Debug)
	if err != nil {
		log.Fatal(err)
	}

	results, err := tile38.Nearby("fleet", 33.462, -112.268, 6000).
		Where("speed", 70, math.MaxInt32).
		Match("truck*").Do()
	if err != nil {
		log.Fatal(err)
	}

	printJSON("nearby:", results)
}

func printJSON(msg string, data interface{}) {
	b, _ := json.Marshal(data)
	fmt.Printf("%s: %s\n", msg, b)
}
