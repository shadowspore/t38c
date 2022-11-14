//go:build ignore
// +build ignore

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

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

	handler := t38c.EventHandlerFunc(func(event *t38c.GeofenceEvent) error {
		b, err := json.Marshal(event)
		if err != nil {
			return fmt.Errorf("marshal event: %w", err)
		}

		fmt.Printf("event: %s\n", b)
		return nil
	})

	if err := tile38.Geofence.Nearby("fleet", 33.462, -112.268, 6000).
		Actions(t38c.Enter, t38c.Exit).
		Do(context.Background(), handler); err != nil {
		log.Fatal(err)
	}
}
