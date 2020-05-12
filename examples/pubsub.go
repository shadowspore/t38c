// +build ignore

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	t38c "github.com/powercake/tile38-client"
)

func main() {
	tile38, err := t38c.New("localhost:9851", t38c.Debug)
	if err != nil {
		log.Fatal(err)
	}

	geofenceRequest := tile38.GeofenceNearby("buses", 33.5123, -112.2693, 200).
		Actions(t38c.Enter, t38c.Exit)

	if err := tile38.SetChan("busstop", geofenceRequest).Do(); err != nil {
		log.Fatal(err)
	}

	handler := func(event *t38c.GeofenceEvent) {
		b, _ := json.Marshal(event)
		fmt.Printf("event: %s\n", b)
	}
	if err := tile38.Subscribe(context.Background(), handler, "busstop"); err != nil {
		log.Fatal(err)
	}
}
