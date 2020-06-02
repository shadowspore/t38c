// +build ignore

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	t38c "github.com/b3q/tile38-client"
)

func main() {
	tile38, err := t38c.New("localhost:9851", t38c.Debug)
	if err != nil {
		log.Fatal(err)
	}

	handler := func(event *t38c.GeofenceEvent) {
		b, _ := json.Marshal(event)
		fmt.Printf("event: %s\n", b)
	}

	if err := tile38.Geofence.Nearby("fleet", 33.462, -112.268, 6000).
		Actions(t38c.Enter, t38c.Exit).
		Do(context.Background(), handler); err != nil {
		log.Fatal(err)
	}
}
