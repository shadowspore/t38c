//go:build ignore
// +build ignore

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"

	"github.com/xjem/t38c"
)

func main() {
	tile38, err := t38c.New(t38c.Config{
		Address: "localhost:9851",
		Debug:   true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer tile38.Close()

	results, err := tile38.Search.Nearby("fleet", 33.462, -112.268, 6000).
		Where("speed", 70, math.MaxInt32).
		Match("truck*").Do(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(results)
	fmt.Printf("nearby: %s\n", b)
}
