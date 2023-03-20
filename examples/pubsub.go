//go:build ignore
// +build ignore

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/sythang/t38c"
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

	geofenceRequest := tile38.Geofence.Nearby("buses", 33.5123, -112.2693, 200).
		Actions(t38c.Enter, t38c.Exit)

	if err := tile38.Channels.SetChan("busstop", geofenceRequest).Do(context.TODO()); err != nil {
		log.Fatal(err)
	}

	handler := t38c.EventHandlerFunc(func(event *t38c.GeofenceEvent) error {
		b, err := json.Marshal(event)
		if err != nil {
			return fmt.Errorf("marshal event: %w", err)
		}

		fmt.Printf("event: %s\n", b)
		return nil
	})

	if err := tile38.Channels.Subscribe(context.Background(), handler, "busstop"); err != nil {
		log.Fatal(err)
	}
}
