package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

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

	geofenceRequest := t38c.GeofenceNearby("fleet", 33.462, -112.268, 6000).
		Actions(t38c.Enter, t38c.Exit)

	events, err := tile38.Fence(context.Background(), geofenceRequest)
	if err != nil {
		log.Fatal(err)
	}

	for event := range events{
		printJSON("event", event)
	}
}
